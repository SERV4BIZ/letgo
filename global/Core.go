package global

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/SERV4BIZ/gfp/collection"
	"github.com/SERV4BIZ/gfp/jsons"
)

// Cache is cache pack data
type Cache struct {
	Header http.Header
	Data   []byte
}

// Request is header pack data
type Request struct {
	Path      string
	SessionID string
	Session   *jsons.JSONObject
	Params    *jsons.JSONObject
	Post      *jsons.JSONObject
	File      *jsons.JSONArray
	MapAPI    *collection.MapKey

	Response http.ResponseWriter
	Request  *http.Request
}

// ReadFile is read file from multipart request
func (me *Request) ReadFile(keyname string) (*jsons.JSONObject, []byte, error) {
	file, header, err := me.Request.FormFile(keyname)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	txtExt := "dat"
	exts := strings.Split(header.Filename, ".")
	if len(exts) >= 2 {
		txtExt = exts[len(exts)-1]
	}

	jsoInfo := jsons.JSONObjectFactory()
	jsoInfo.PutString("txt_name", header.Filename)
	jsoInfo.PutInt("int_size", int(header.Size))
	jsoInfo.PutString("txt_ext", txtExt)

	buffer, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, nil, err
	}

	return jsoInfo, buffer, nil
}

// AddAPI is add api for head pack
func (me *Request) AddAPI(path string, apifunc interface{}) {
	key := strings.ToLower(path)
	me.MapAPI.Put(key, apifunc)
}

// GetAPI is get api for head pack
func (me *Request) GetAPI(path string) func(rep *Request) *jsons.JSONObject {
	key := strings.ToLower(path)
	if me.MapAPI.ContainsKey(key) {
		return me.MapAPI.Get(key).(func(rep *Request) *jsons.JSONObject)
	}
	return nil
}

// SessionName is session name
const SessionName string = "LETGO_SESSION_ID"

// DS is separator of directory
const DS string = "/"

// MaxMemoryMultipart is Max memory 100MB
const MaxMemoryMultipart int64 = 100 * 1024 * 1024

// MaxSessionExpire is Max expire of session
const MaxSessionExpire time.Duration = 24 * 365 * time.Hour

// IsCacheWork is enable cache work
var IsCacheWork bool = false

// MutexListProtect is security directory
var MutexListProtect sync.RWMutex

// ListProtect is list security directory
var ListProtect []string

// MutexMapCache is mutex lock for MapCache
var MutexMapCache sync.RWMutex

// MapCache is map for Cache
var MapCache map[string]*Cache = make(map[string]*Cache)

// MutexMapMimeType is mutex lock for MapMimeType
var MutexMapMimeType sync.RWMutex

// MapMimeType is map for mime type data
var MapMimeType map[string]string = make(map[string]string)

// MutexMemCache is mutex lock for MemCache
var MutexMemCache sync.RWMutex

// MemCache is memory cache
var MemCache *jsons.JSONObject = jsons.JSONObjectFactory()

// MutexState is mutex lock MemoryState
var MutexState sync.RWMutex

// MemoryState is memory state
var MemoryState int = 0

// LoadState is load state
var LoadState int = 0

// CountState is count state per second
var CountState int = 0
