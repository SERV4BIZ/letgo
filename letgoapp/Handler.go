package letgoapp

import (
	"fmt"
	"time"

	"github.com/SERV4BIZ/gfp/hash"
	"github.com/SERV4BIZ/gfp/uuid"
	"github.com/SERV4BIZ/letgo/global"
	"github.com/SERV4BIZ/letgo/utility"
)

// RegisterAPIHandler is variable api all method
var RegisterAPIHandler interface{} = func(rep *global.Request) {}

// CustomSessionIDHandler is custom session id for all request
var CustomSessionIDHandler interface{} = func(rep *global.Request) {
	txtUUID, errUUID := uuid.NewV4()
	if errUUID != nil {
		txtUUID = hash.SHA256([]byte(fmt.Sprint(time.Now().Unix())))
	}
	rep.SessionID = txtUUID
}

// LoadSessionHandler is start session of all request
var LoadSessionHandler interface{} = func(rep *global.Request) {
	filename := fmt.Sprint(rep.SessionID, ".session")
	pathfile := fmt.Sprint(utility.GetAppDir(), global.DS, "sessions", global.DS, filename)
	rep.Session.FromFile(pathfile)
}

// SaveSessionHandler is flush session of all request
var SaveSessionHandler interface{} = func(rep *global.Request) {
	filename := fmt.Sprint(rep.SessionID, ".session")
	pathfile := fmt.Sprint(utility.GetAppDir(), global.DS, "sessions", global.DS, filename)
	rep.Session.ToFile(pathfile)
}
