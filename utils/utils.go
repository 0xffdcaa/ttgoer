package utils

import (
	"bytes"
	"runtime"
)

func Contains[T comparable](slice []T, value T) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func GetGoroutineName() string {
	buf := make([]byte, 64)
	buf = buf[:runtime.Stack(buf, false)]
	fields := bytes.Fields(buf)
	return string(fields[1])
}
