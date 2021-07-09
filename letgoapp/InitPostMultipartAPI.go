package letsgoapp

import "github.com/SERV4BIZ/letsgo/global"

// InitPostMultipartAPI is load post multipart form for api
func InitPostMultipartAPI(rep *global.Request) {
	err := rep.Request.ParseMultipartForm(int64(global.MaxRead))
	if err == nil {
		for key, val := range rep.Request.PostForm {
			if len(val) > 1 {
				rep.POST.GetObjectData().Put(key, val)
			} else {
				rep.POST.GetObjectData().Put(key, val[0])
			}
		}

		for key := range rep.Request.MultipartForm.File {
			rep.FILE.PutString(key)
		}
	}
}
