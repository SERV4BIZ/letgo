package main

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/letgo/letgoapp"
	"github.com/SERV4BIZ/letgo/utility"
)

// APIDIR is path of folder apis
var APIDIR string = "apis"

// ScanDir is listing file in folder
func ScanDir(folder string) *jsons.JSONArray {
	jsaResult := jsons.JSONArrayFactory()
	filedirs, errScan := files.ScanDir(folder)
	if errScan == nil {
		for _, val := range filedirs {
			if val != ".DS_Store" && val != "code.js" && val != "view.html" && val != "style.css" {
				pathFile := fmt.Sprint(folder, "/", val)
				if files.IsFile(pathFile) {
					jsaResult.PutString(pathFile)
				} else {
					jsaList := ScanDir(pathFile)
					for i := 0; i < jsaList.Length(); i++ {
						jsaResult.PutString(jsaList.GetString(i))
					}
				}
			}
		}
	}

	return jsaResult
}

// APIListing is listing api all in folder
func APIListing(folder string) *jsons.JSONArray {
	pathdir := fmt.Sprint(utility.GetAppDir(), "/", folder)

	jsaNList := jsons.JSONArrayFactory()
	jsaList := ScanDir(pathdir)
	for i := 0; i < jsaList.Length(); i++ {
		jsaNList.PutString(strings.TrimPrefix(fmt.Sprint("/", strings.Trim(strings.Replace(jsaList.GetString(i), utility.GetAppDir(), "", 1), "/")), fmt.Sprint("/", APIDIR, "/")))
	}

	return jsaNList
}

func main() {
	ProjectName := "LetGo Compiler"
	ProjectVersion := "1.0.0"
	CompanyName := "SERV4BIZ CO.,LTD."

	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint(ProjectName, " Version ", ProjectVersion))
	fmt.Println(fmt.Sprint("Copyright © 2020 ", CompanyName, " All Rights Reserved."))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint("Directory : ", utility.GetAppDir()))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	jsaListing := APIListing(APIDIR)

	jsaPathAPI := jsons.JSONArrayFactory()
	mapPathImport := make(map[string]string)
	for i := 0; i < jsaListing.Length(); i++ {
		fmt.Println(jsaListing.GetString(i))
		txtPathFile := jsaListing.GetString(i)

		arrPaths := strings.Split(txtPathFile, "/")
		txtObjectName := ""
		txtImportPath := ""
		txtImportName := ""
		for _, val := range arrPaths {
			if strings.HasSuffix(val, ".go") {
				txtObjectName = fmt.Sprint(txtObjectName, ".", strings.TrimSuffix(val, ".go"))
				txtObjectName = strings.TrimPrefix(txtObjectName, "_")
				break
			} else {
				txtObjectName = fmt.Sprint(txtObjectName, "_", val)

				txtImportPath = fmt.Sprint(txtImportPath, "/", val)
				txtImportName = fmt.Sprint(txtImportName, "_", val)
			}
		}
		txtImportPath = strings.TrimPrefix(txtImportPath, "/")
		txtImportName = strings.TrimPrefix(txtImportName, "_")
		mapPathImport[txtImportPath] = txtImportName

		txtPathAPI := strings.TrimSuffix(txtPathFile, ".go")
		jsaPathAPI.PutString(fmt.Sprint("rep.AddAPI(\"", txtPathAPI, "\",", txtObjectName, ")"))
	}

	txtImportBuffer := ""
	for keyName := range mapPathImport {
		txtImportBuffer = fmt.Sprint(txtImportBuffer, mapPathImport[keyName], " \"./", APIDIR, "/", keyName, "\"\n")
	}

	txtAddAPIBuffer := ""
	for i := 0; i < jsaPathAPI.Length(); i++ {
		txtAddAPIBuffer = fmt.Sprint(txtAddAPIBuffer, jsaPathAPI.GetString(i), "\n")
	}

	codeBuffer := `
	package main

	import(
		"github.com/SERV4BIZ/letgo/global"
		"github.com/SERV4BIZ/letgo/letgoapp"

		` + txtImportBuffer + `
	)

	func LetGoAPI() {
		letgoapp.RegisterAPIHandler = func(rep *global.Request) {
			` + txtAddAPIBuffer + `
		}
	}`

	pathFile := fmt.Sprint(utility.GetAppDir(), "/apis.go")
	files.WriteFile(pathFile, []byte(codeBuffer))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint(ProjectName, " Finished"))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	LetGoAPI()
	letgoapp.Listen(0)
}
