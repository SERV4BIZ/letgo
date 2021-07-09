package letsgoapp

import (
	"fmt"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/SERV4BIZ/gfp/collection"
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/letsgo/global"
	"github.com/SERV4BIZ/letsgo/utility"
)

// Listen is begin work
func Listen(port int) {
	jsoConfig, errConfig := utility.GetConfig("letsgo")
	if errConfig != nil {
		panic(fmt.Sprint("Can not load config letsgo.json file [ ", errConfig, " ]"))
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
		global.MapMimeType[extName] = strings.ToLower(jsoMimetype.GetString(extName))
	}

	pathApp := utility.GetAppDir()
	jsaProtect := jsoConfig.GetArray("jsa_protect")
	for i := 0; i < jsaProtect.Length(); i++ {
		txtItem := jsaProtect.GetString(i)
		pathProtect := fmt.Sprint(pathApp, "/", strings.Trim(txtItem, "/"))
		global.ListProtect = append(global.ListProtect, pathProtect)
	}

	global.MaxRead = jsoConfig.GetInt("int_maxread")
	if global.MaxRead <= 0 {
		// Default max reader is 1024MB or 1GB
		global.MaxRead = 1024 * 1024 * 1024
	}

	maxSession := jsoConfig.GetInt("int_maxsession")
	if maxSession > 0 {
		global.MaxSession = time.Duration(maxSession) * time.Minute
	} else {
		// Default max session 30 minute
		global.MaxSession = 30 * time.Minute
	}

	// Load and Memory Monitor
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

	// Force GC to clear up
	go func() {
		for {
			<-time.After(time.Hour)
			runtime.GC()
		}
	}()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		global.MutexState.Lock()
		global.CountState++
		global.MutexState.Unlock()

		r.Body = http.MaxBytesReader(w, r.Body, int64(global.MaxRead))
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		rep := new(global.Request)
		rep.Path = r.URL.Path
		rep.SessionID = ""

		rep.SESSION = jsons.ObjectNew()
		rep.GET = jsons.ObjectNew()
		rep.POST = jsons.ObjectNew()
		rep.FILE = jsons.ArrayNew()

		rep.MapAPI = collection.MapKeyFactory()
		rep.Response = w
		rep.Request = r
		Process(rep)
	})

	http.ListenAndServe(fmt.Sprint("0.0.0.0:", port), nil)
}
