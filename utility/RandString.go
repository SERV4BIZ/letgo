package utility

import (
	"fmt"
	"math/rand"
	"time"
)

// RandString is random string of level
func RandString(n int, level int) string {
	letterBytes := "0123456789"
	if level == 1 {
		letterBytes = "0123456789abcdefghijklmnopqrstuvwxyz"
	} else if level >= 2 {
		letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	result := ""
	for i := 0; i < n; i++ {
		rand.Seed(time.Now().UnixNano())
		result = fmt.Sprint(result, string(letterBytes[rand.Intn(len(letterBytes))]))
	}
	return result
}
