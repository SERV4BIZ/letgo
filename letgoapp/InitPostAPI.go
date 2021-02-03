package letgoapp

import "github.com/SERV4BIZ/letgo/global"

// InitPostAPI is load post from for api
func InitPostAPI(rep *global.Request) {
	err := rep.Request.ParseForm()
	if err == nil {
		for key, val := range rep.Request.PostForm {
			if len(val) > 1 {
				rep.POST.GetObjectData().Put(key, val)
			} else {
				rep.POST.GetObjectData().Put(key, val[0])
			}
		}
	}
}
