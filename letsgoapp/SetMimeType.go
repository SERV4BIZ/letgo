package letsgoapp

import (
	"github.com/SERV4BIZ/letsgo/global"
	"github.com/SERV4BIZ/letsgo/utility"
)

// SetMimeType is set header type for response
func SetMimeType(rep *global.Request) {
	ext := GetPathExt(rep.Path)
	mt, err := utility.MimeType(ext)
	if err == nil {
		rep.Response.Header().Set("Content-Type", mt)
	}
}
