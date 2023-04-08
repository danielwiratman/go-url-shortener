package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomString(length int) string {
	rand := rand.New(rand.NewSource(time.Now().UnixNano())) // convention
	charset := "abcdefghijklmnopqrstuvwxyz"
	byteSeq := make([]byte, length)
	for i := 0; i < length; i++ {
		c := charset[rand.Intn(len(charset))]
		byteSeq[i] = c
	}
	return string(byteSeq)
}