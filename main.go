package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	const (
		securityparam int = 1000000
	)
	var (
		publicKey int
		secretKey int
	)
	fmt.Println("hello, world", securityparam)
	rsa.GenerateKey(rand.Int(securityparam))
}

func GenerateKey(securityparam int) int {
	rsa.GenerateKey()
}
