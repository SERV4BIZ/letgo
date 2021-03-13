package letgoapp

import (
	"net/http"
	"time"

	"github.com/SERV4BIZ/letgo/global"
)

// InitSession is load session first
func InitSession(rep *global.Request) string {
	cookie, err := rep.Request.Cookie(global.SessionName)
	if err != nil {
		CustomSessionIDHandler.(func(rep *global.Request))(rep)
		expiration := time.Now().Add(global.MaxSession)
		cookie = &http.Cookie{Name: global.SessionName, Value: rep.SessionID, Expires: expiration, Path: "/"}
		http.SetCookie(rep.Response, cookie)
	}
	rep.SessionID = cookie.Value
	LoadSessionHandler.(func(rep *global.Request))(rep)
	return rep.SessionID
}
