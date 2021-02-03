package letgoapp

import "./global"

// FlushSession is save session to file
func FlushSession(rep *global.Request) {
	SaveSessionHandler.(func(rep *global.Request))(rep)
}
