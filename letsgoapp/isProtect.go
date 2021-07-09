package letsgoapp

import (
	"strings"

	"github.com/SERV4BIZ/letsgo/global"
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
