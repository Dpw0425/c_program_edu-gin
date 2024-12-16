package generator

import (
	"math/rand"
	"time"
)

func Random(length int) string {
	var result = make([]byte, length)
	charSet := []byte("0123456789abcdefghijklmnopqrstuvwxyz")

	b := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := range result {
		result[i] = charSet[b.Intn(len(charSet))]
	}

	return string(result)
}
