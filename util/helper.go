package util

import (
	"math/rand"
	"strings"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func StringArrayToString(inputArr []string) string {
	return strings.Join(inputArr, ",")
}

func StringToStringArray(inputStr string) []string {
	return strings.Split(inputStr, ",")
}
