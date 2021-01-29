package networks

import (
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/letgo/global"
)

// Ping is check network
func Ping(rep *global.Request) *jsons.JSONObject {
	jsoResult := jsons.JSONObjectFactory()
	jsoResult.PutInt("status", 1)

	global.MutexState.RLock()
	jsoData := jsons.JSONObjectFactory()
	jsoData.PutInt("int_memory", global.MemoryState)
	jsoData.PutInt("int_load", global.LoadState)
	jsoResult.PutObject("jso_data", jsoData)
	global.MutexState.RUnlock()

	return jsoResult
}
