package letsgoapp

import "github.com/SERV4BIZ/letsgo/global"

// FlushSession is save session to file
func FlushSession(rep *global.Request) {
	SaveSessionHandler.(func(rep *global.Request))(rep)
}
