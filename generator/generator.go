package generator

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz0123456789"

var src = rand.NewSource(time.Now().UnixNano())

func RandStringBytes() string {
	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[src.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
