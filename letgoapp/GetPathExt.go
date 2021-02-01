package letgoapp

import (
	"strings"
)

// GetPathExt is get ext file of path
func GetPathExt(path string) string {
	paths := strings.Split(path, ".")
	return strings.ToLower(paths[len(paths)-1])
}
