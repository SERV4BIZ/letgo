package letgo

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/SERV4BIZ/gfp/files"
	"github.com/SERV4BIZ/letgo/global"
	"github.com/SERV4BIZ/letgo/utility"
)

// Process is main processing
func Process(rep *global.Request) {
	pathfile := fmt.Sprint(utility.GetAppDir(), rep.Path)
	if strings.TrimSpace(rep.Path) == "/" || isProtect(pathfile) {
		pathfile = fmt.Sprint(pathfile, "index.html")
		rep.Response.Header().Set("Content-Type", "text/html")
	}

	if global.IsCacheWork {
		global.MutexMapCache.RLock()
		cache, ok := global.MapCache[pathfile]
		global.MutexMapCache.RUnlock()

		if ok {
			SetMimeType(rep)
			global.MutexMapCache.RLock()
			for key, value := range cache.Header {
				rep.Response.Header().Add(key, value[0])
			}
			global.MutexMapCache.RUnlock()

			rep.Response.Write(cache.Data)
		} else {
			if files.ExistFile(pathfile) {
				if files.IsFile(pathfile) {
					buff, errBuff := files.ReadFile(pathfile)
					if errBuff == nil {
						cache = new(global.Cache)
						buff = []byte(TagMap(string(buff)))
						cache.Data = buff

						SetMimeType(rep)
						rep.Response.Write(cache.Data)

						cache.Header = http.Header{}
						for key, value := range rep.Response.Header() {
							cache.Header.Add(key, value[0])
						}

						global.MutexMapCache.Lock()
						global.MapCache[pathfile] = cache
						global.MutexMapCache.Unlock()
					} else {
						fmt.Fprint(rep.Response, fmt.Sprint("URL Path ", rep.Path, " Error [", errBuff, "]"))
					}
				} else {
					RenderAPI(rep)
				}
			} else {
				RenderAPI(rep)
			}
		}
	} else {
		if files.ExistFile(pathfile) {
			if files.IsFile(pathfile) {
				buff, errBuff := files.ReadFile(pathfile)
				if errBuff == nil {
					SetMimeType(rep)
					buff = []byte(TagMap(string(buff)))
					rep.Response.Write(buff)
				} else {
					fmt.Fprint(rep.Response, fmt.Sprint("URL Path ", rep.Path, " Error [", errBuff, "]"))
				}
			} else {
				RenderAPI(rep)
			}
		} else {
			RenderAPI(rep)
		}
	}
}
