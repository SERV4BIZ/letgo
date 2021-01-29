package utility

import (
	"errors"

	"github.com/SERV4BIZ/letgo/global"
)

// MimeType is set header type for response
func MimeType(txtExt string) (string, error) {
	global.MutexMapMimeType.RLock()
	mimetype, ok := global.MapMimeType[txtExt]
	global.MutexMapMimeType.RUnlock()

	if !ok {
		return "", errors.New("Mime type not support")
	}

	return mimetype, nil
}
