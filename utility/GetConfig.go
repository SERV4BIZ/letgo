package utility

import (
	"fmt"
	"sync"

	"github.com/SERV4BIZ/gfp/collection"
	"github.com/SERV4BIZ/gfp/jsons"
	"github.com/SERV4BIZ/letgo/global"
)

// MutexMapConfig is mutex lock for MapConfig
var MutexMapConfig sync.RWMutex

// MapConfig is map for config object
var MapConfig *collection.MapKey = collection.MapKeyFactory()

// GetConfig is get config json object
func GetConfig(name string) (*jsons.JSONObject, error) {
	var jsoConfig *jsons.JSONObject = nil
	var errConfig error = nil

	MutexMapConfig.RLock()
	if MapConfig.HasKey(name) {
		objConfig := MapConfig.Get(name)
		if objConfig != nil {
			jsoConfig, errConfig = objConfig.(*jsons.JSONObject).Copy()
		}
	}
	MutexMapConfig.RUnlock()

		return nil, errConfig
	}

	if jsoConfig != nil {
		return jsoConfig, nil
	}

	MutexMapConfig.Lock()
	pathfile := fmt.Sprint(GetAppDir(), global.DS, "configs", global.DS, name, ".json")
	jsoConfig, errConfig = jsons.JSONObjectFromFile(pathfile)
	if errConfig == nil {
		MapConfig.Put(name, jsoConfig)
	}
	MutexMapConfig.Unlock()

	if errConfig != nil {
		return nil, errConfig
	}
	return jsoConfig.Copy()
}
