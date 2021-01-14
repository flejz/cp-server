package util

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func Hash(value string, salt string) string {
	bytes := sha256.Sum256([]byte(fmt.Sprintf("%s.%s", value, salt)))
	return string(bytes[:])
}

func Salt() string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder

	charSet := []rune("abcdedfghijklmnopqrstABCDEFGHIJKLMNOP0123456789")
	length := 20
	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteRune(randomChar)
	}
	return output.String()
}
