package letgoapp

import "github.com/SERV4BIZ/letgo/global"

// InitPostMultipartAPI is load post multipart form for api
func InitPostMultipartAPI(rep *global.Request) {
	err := rep.Request.ParseMultipartForm(global.MaxMemoryMultipart)
	if err == nil {
		for key, val := range rep.Request.PostForm {
			if len(val) > 1 {
				rep.Post.GetObjectData().Put(key, val)
			} else {
				rep.Post.GetObjectData().Put(key, val[0])
			}
		}

		for key := range rep.Request.MultipartForm.File {
			rep.File.PutString(key)
		}
	}
}
