package letgoapp

import (
	"fmt"
	"strings"

	"github.com/SERV4BIZ/letgo/global"
)

// InitParamsAPI is load parameter for api path
func InitParamsAPI(rep *global.Request) {
	path := strings.TrimSuffix(strings.ToLower(rep.Path), ".html")
	params := strings.Split(path, global.DS)

	pathapi := ""
	count := 0
	for index, val := range params {
		if index >= 2 {
			pathapi = fmt.Sprint(pathapi, "/", val)
			rep.Params.PutString(fmt.Sprint("var", count), val)
			count++
		}
	}

	pathapi = strings.Trim(pathapi, "/")
	rep.Params.PutString("path", pathapi)

	// Check parameter
	for key, val := range rep.Request.URL.Query() {
		if len(val) > 1 {
			rep.Params.GetObjectData().Put(key, val)
		} else {
			rep.Params.GetObjectData().Put(key, val[0])
		}
	}
}
