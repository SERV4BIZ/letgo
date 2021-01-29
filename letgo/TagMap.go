package letgo

import (
	"fmt"
	"strings"
	"time"
)

// TagMap is change tag in file to system info
func TagMap(buffer string) string {
	nbuff := buffer
	nbuff = strings.ReplaceAll(nbuff, "{{STAMP}}", fmt.Sprint(time.Now().Unix()))
	return nbuff
}
