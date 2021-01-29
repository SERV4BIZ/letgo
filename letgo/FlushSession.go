package letgo

import "github.com/SERV4BIZ/letgo/global"

// FlushSession is save session to file
func FlushSession(rep *global.Request) {
	SaveSessionHandler.(func(rep *global.Request))(rep)
}
