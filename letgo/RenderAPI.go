package letgo

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/letgo/global"
	"github.com/SERV4BIZ/letgo/utility"
)

// APINotFound is render api not found
func APINotFound(rep *global.Request) {
	jsoResult := jsons.JSONObjectFactory()
	jsoResult.PutInt("status", 0)
	jsoResult.PutString("txt_msg", "API NOT FOUND")
	fmt.Fprint(rep.Response, jsoResult.ToString())
}

// RenderAPI is render for api response
func RenderAPI(rep *global.Request) {
	path := strings.TrimSuffix(strings.ToLower(rep.Path), ".html")
	params := strings.Split(path, global.DS)
	if strings.ToLower(params[1]) == "api" {
		if rep.Request.Method == "POST" {
			InitSession(rep)
			InitParamsAPI(rep)
			InitPostAPI(rep)
			InitPostMultipartAPI(rep)
			regisapi := RegisterAPIHandler.(func(rep *global.Request))
			regisapi(rep)

			path := rep.Params.GetString("path")
			api := rep.GetAPI(path)
			if api != nil {
				rep.Response.Header().Set("Content-Type", "application/json")
				jsoResult := api(rep)
				if jsoResult != nil {
					fmt.Fprint(rep.Response, jsoResult.ToString())
				}
			} else {
				APINotFound(rep)
			}
			FlushSession(rep)
		} else if rep.Request.Method == "GET" {
			InitParamsAPI(rep)
			regisapi := RegisterAPIHandler.(func(rep *global.Request))
			regisapi(rep)

			path := rep.Params.GetString("path")
			api := rep.GetAPI(path)
			if api != nil {
				rep.Response.Header().Set("Content-Type", "application/json")
				jsoResult := api(rep)
				if jsoResult != nil {
					fmt.Fprint(rep.Response, jsoResult.ToString())
				}
			} else {
				APINotFound(rep)
			}
		} else {
			APINotFound(rep)
		}
	} else {
		pathfile := fmt.Sprint(utility.GetAppDir(), "/index.html")
		buff, errBuff := files.ReadFile(pathfile)
		if errBuff == nil {
			rep.Response.Header().Set("Content-Type", "text/html")
			buff = []byte(TagMap(string(buff)))
			rep.Response.Write(buff)
		}
	}
}
