package letgo

import (
	"strings"

	"github.com/SERV4BIZ/letgo/global"
)

func isProtect(pathFile string) bool {
	global.MutexListProtect.RLock()
	defer global.MutexListProtect.RUnlock()

	for _, protectItem := range global.ListProtect {
		if strings.HasPrefix(pathFile, protectItem) {
			return true
		}
	}

	return false
}
