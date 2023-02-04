/*
 * @Author: youlingdada youlingdada@163.com
 * @Date: 2022-07-08 09:36:15
 * @LastEditors: youlingdada youlingdada@163.com
 * @LastEditTime: 2022-07-08 09:52:31
 * @FilePath: \street-stall\utils\crypt.go
 * @Description: 加密工具
 * QQ:3367758294
 * Copyright (c) 2022 by Youlingdada, All Rights Reserved.
 */

package utils

import "golang.org/x/crypto/bcrypt"

/**
 * @description: 加密
 * @param {string} plain 明文
 * @return {*} 加密后的密文
 */
func CryptEncode(plain string) (*string, error) {
	cipher, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	res := string(cipher)
	return &res, nil
}

/**
 * @description: 解密
 * @param {string} plain 明文
 * @param {string} cipher 密文
 * @return {*} 明文是否匹配密文
 */
func CryptMatch(plain, cipher string) error {
	return bcrypt.CompareHashAndPassword([]byte(cipher), []byte(plain))
}
