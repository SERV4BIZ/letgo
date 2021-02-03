package letgoapp

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/SERV4BIZ/gfp/collection"
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/letgo/global"
	"github.com/SERV4BIZ/letgo/utility"
)

// Listen is begin work
func Listen(port int) {
	jsoConfig, errConfig := utility.GetConfig("letgo")
	if errConfig != nil {
		panic(fmt.Sprint("Can not load config letgo.json file [ ", errConfig, " ]"))
	}

	if port <= 0 {
		port = jsoConfig.GetInt("int_port")
	}

	global.IsCacheWork = jsoConfig.GetInt("int_cache") > 0

	jsoMimetype, errMimetype := utility.GetConfig("mimetype")
	if errMimetype != nil {
		panic(fmt.Sprint("Can not load config mimetype.json file [ ", errMimetype, " ]"))
	}
	exts := jsoMimetype.GetKeys()
	for _, extName := range exts {
		global.MapMimeType[extName] = jsoMimetype.GetString(extName)
	}

	pathApp := utility.GetAppDir()
	jsaProtect := jsoConfig.GetArray("jsa_protect")
	for i := 0; i < jsaProtect.Length(); i++ {
		txtItem := jsaProtect.GetString(i)
		pathProtect := fmt.Sprint(pathApp, "/", strings.Trim(txtItem, "/"))
		global.ListProtect = append(global.ListProtect, pathProtect)
	}

	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			global.MutexState.Lock()
			global.MemoryState = int(utility.NumberByteToMb(m.Sys))
			global.LoadState = global.CountState
			global.CountState = 0
			global.MutexState.Unlock()

			<-time.After(time.Second)
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		global.MutexState.Lock()
		global.CountState++
		global.MutexState.Unlock()

		r.Body = http.MaxBytesReader(w, r.Body, global.MaxMemoryMultipart)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		rep := new(global.Request)
		rep.Path = r.URL.Path
		rep.SessionID = ""

		rep.SESSION = jsons.ObjectNew()
		rep.GET = jsons.ObjectNew()
		rep.POST = jsons.ObjectNew()
		rep.FILE = jsons.ObjectNew()

		rep.MapAPI = collection.MapKeyFactory()
		rep.Response = w
		rep.Request = r
		Process(rep)
	})

	http.ListenAndServe(fmt.Sprint("0.0.0.0:", port), nil)
}
