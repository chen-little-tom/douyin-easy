package user_token

import (
	"douyin-easy/utils/ras_util"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"strings"
)

var tokenPrefix *string = nil

// GetTokenPrefix 单例模式，获取一个token的前缀
func GetTokenPrefix() (string, error) {
	if tokenPrefix != nil {
		return *tokenPrefix, nil
	}
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	var s string
	s = newUUID.String()
	tokenPrefix = &s
	return s, err
}

// GetUserToken 获取user user_token
func GetUserToken(userId uint64) (string, error) {
	prefix, err := GetTokenPrefix()
	if err != nil {
		return "", err
	}
	uId := strconv.FormatUint(userId, 10)
	rsaEncode, err := ras_util.RsaEncode(uId)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s-%s", prefix, *rsaEncode), nil
}

// GetUserIdByToken 根据token 获取userId
// user_token 令牌
func GetUserIdByToken(token string) (uint64, error) {
	// 处理rsa传输过程中 + -> " "
	strings.ReplaceAll(token, " ", "+")
	index := strings.LastIndex(token, "-")
	s := token[index+1:]
	decode, err := ras_util.RSADecode(s)
	if err != nil {
		return 0, err
	}
	userId, err := strconv.ParseUint(*decode, 10, 64)
	if err != nil {
		return 0, err
	}
	return userId, nil
}
