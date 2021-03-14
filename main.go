package main

import (
	"fmt"

	"strings"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
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
			pathFile := fmt.Sprint(folder, "/", val)
			if files.IsFile(pathFile) {
				if strings.HasSuffix(strings.ToLower(val), ".go") {
					jsaResult.PutString(pathFile)
				}
			} else {
				jsaList := ScanDir(pathFile)
				for i := 0; i < jsaList.Length(); i++ {
					jsaResult.PutString(jsaList.GetString(i))
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

	pathModfile := fmt.Sprint(utility.GetAppDir(), "/go.mod")
	if !files.ExistFile(pathModfile) {
		panic("Not found go.mod file")
	}
	modBuff, errMod := files.ReadFile(pathModfile)
	if errMod != nil {
		panic(errMod)
	}
	txtModBuff := string(modBuff)
	moduleName := strings.ReplaceAll(strings.Split(txtModBuff, "\n")[0], "module ", "")

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

	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint("Initial Module All API"))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	txtImportBuffer := ""
	for keyName := range mapPathImport {
		pathPackage := fmt.Sprint(moduleName, "/", APIDIR, "/", keyName)
		txtImportBuffer = fmt.Sprint(txtImportBuffer, mapPathImport[keyName], " \"", pathPackage, "\"\n")

		// Init module and tidy and get
		fmt.Println(pathPackage)
		// init
		//cmd := exec.Command("go", "mod", "init", pathPackage)
		//cmd.Dir = fmt.Sprint(APIDIR, "/", keyName)
		//cmd.Run()

		// tidy
		/*cmd = exec.Command("go", "mod", "tidy")
		cmd.Dir = fmt.Sprint(APIDIR, "/", keyName)
		cmd.Run()*/

		// get
		/*cmd = exec.Command("go", "get", pathPackage)
		cmd.Run()*/

		// replace to local
		//cmd = exec.Command("go", "mod", "edit", "-replace", fmt.Sprint(pathPackage, "=", "./apis/", keyName))
		//cmd.Dir = fmt.Sprint(APIDIR)
		//cmd.Run()
		// end
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

	// project base module tidy
	//cmd := exec.Command("go", "mod", "tidy")
	//cmd.Dir = fmt.Sprint(utility.GetAppDir())
	//cmd.Run()

	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")
	fmt.Println(fmt.Sprint(ProjectName, " Finished"))
	fmt.Println("* * * * * * * * * * * * * * * * * * * * * * * * * * * * * *")

	//LetGoAPI()
	//letgoapp.Listen(0)
}
