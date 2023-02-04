package ras_util

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"testing"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey

func TestMain(m *testing.M) {
	private, _ := rsa.GenerateKey(rand.Reader, 2048)
	public := private.PublicKey
	privateKey = private
	publicKey = &public

	key, _ := GetRSAPublicKey(publicKey)
	fmt.Println(key)
	rsaPrivateKey := GetRSAPrivateKey(privateKey)
	fmt.Println(rsaPrivateKey)

	encode, err := RsaEncode("123456")
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(*encode)
	decode, err := RSADecode(*encode)
	fmt.Println(*decode)
}
