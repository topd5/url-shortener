package main

import (
	"math/rand"
)

func GenerateRandomString() string {
	// Can't use base64.StdEncoding.EncodeToString here because of "/" и "=" и "+" characters

	baseStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	randomString := ""

	for i := 0; i < 7; i++ {
		randomString += string([]rune(baseStr)[rand.Intn(len(baseStr))])
	}

	return randomString
}
