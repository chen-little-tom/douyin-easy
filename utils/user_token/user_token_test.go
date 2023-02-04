package user_token

import (
	"fmt"
	"testing"
)

func TestGetUserIdByToken(t *testing.T) {
	token, _ := GetUserToken(234)
	userId, _ := GetUserIdByToken(token)
	fmt.Println(userId)
}

func TestGetUserToken(t *testing.T) {
	token, _ := GetUserToken(234)
	fmt.Println(token)
}

func TestGetTokenPrefix(t *testing.T) {
	prefix, _ := GetTokenPrefix()
	fmt.Println(prefix)
}
