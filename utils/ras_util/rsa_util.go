package ras_util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

var publicKeyPem string = "-----BEGIN RSA PUBLIC KEY-----\n" +
	"MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEApyFsZeqjjXF17xgizilj\n" +
	"8CITj9hGllbjY6eGEp1reSqVE28vgwZDnDqrBcF5RLsHKPW0y0p4cM2H5Y5tf6j5\n" +
	"8DA6X8XyE/U19sU1rvAY8CtighSVA0jZJzru/S21jIqcfPiwL2WZ0BGlP16KRo0i\n" +
	"yer+MziQYUEaFiv5+y5u1qP18oj2mOma6xWDyecQQS50YvOigDaTY346Iybg8nna\n" +
	"pskrHrIT/kC77ksqK84NI7UiCXbk2588khpYcGyL+T9VGckAV0yNNBG0olUj3xJH\n" +
	"TfIXXGYZZFYJYmz2fjW1lMnir1+Wv7Fo3yRW38cB7WVhSHmWtPgzgSTjjWghbpEz\n" +
	"gQIDAQAB\n" +
	"-----END RSA PUBLIC KEY-----"
var privateKeyPem string = "-----BEGIN RSA PRIVATE KEY-----\n" +
	"MIIEpQIBAAKCAQEApyFsZeqjjXF17xgizilj8CITj9hGllbjY6eGEp1reSqVE28v\n" +
	"gwZDnDqrBcF5RLsHKPW0y0p4cM2H5Y5tf6j58DA6X8XyE/U19sU1rvAY8CtighSV\n" +
	"A0jZJzru/S21jIqcfPiwL2WZ0BGlP16KRo0iyer+MziQYUEaFiv5+y5u1qP18oj2\n" +
	"mOma6xWDyecQQS50YvOigDaTY346Iybg8nnapskrHrIT/kC77ksqK84NI7UiCXbk\n" +
	"2588khpYcGyL+T9VGckAV0yNNBG0olUj3xJHTfIXXGYZZFYJYmz2fjW1lMnir1+W\n" +
	"v7Fo3yRW38cB7WVhSHmWtPgzgSTjjWghbpEzgQIDAQABAoIBAQClmbeXkJP4LYlG\n" +
	"nGq4TkjJAmrRp+HhSzzKXI67WSHJkEjgVdYFBKrvXS5iaJ8pXAwvKK64lBSc6PoT\n" +
	"mjmjOaJVvOPKTA+eeOS8nBGiQiLjrolQDObNt1v7xChT0vvFwv5l0eB5RjnO8f0D\n" +
	"ukBWwbfLzL2NsSlWRDR97ZsnqI2IvJlnHJvs675ivbDCuFh5mV3BPgSXDLZ8kigq\n" +
	"DObnV11nCTQwQDYtmysw6a5BjD6wWcFqB+8+fZ1eIfXikrfA0LVyeWNawxwzpmZF\n" +
	"DtBJSUQDZEqBg9Ez1iOW0sp/OWEu2ySuwGbBWHPeiBzYBL/fQ7v/HJsFShfmOkT1\n" +
	"1MzOHsXlAoGBANZBM6Y8zzcL9zN68Euwtpn/h2yO7Hy+JQXcYyku4AtKXzJy6Ap2\n" +
	"p4l2dk9cvk1MHpaD6pSyGy43yh4MtnPcB7kg079V1UZC0zo7UYD+npvAbN4OWxwu\n" +
	"92qTI2dVMVujgzIvkixWX+xYlj+26BeLJpvF3gt7IXCVQejL51+8EfiHAoGBAMex\n" +
	"uDS24yRK3lsn71b9IkXKHbl29o+rPVha+fGw5cdfUiq/gcYDznpW6vhjCn9RAzcH\n" +
	"Dv+rwO7wD0ho+9P0MyxUjLqa1q41G6sIWAoFskO25bw6lB9coSQICvpo8xEXlP4t\n" +
	"qiRer2YM+6YiRuTJIee0xlMx1u4gqUuYji97dN23AoGBAJ2w0X3Zpb1jDicSpcdf\n" +
	"oZXIb3jj3ISXQhKFOWp3OKi2rUgpMEV7PSW8TaltnhawHpczMjUMvVqP7y+ctbCi\n" +
	"UgDce1yIpPcYefSS8hLZ1AzYXIg20rH0k18aOmV0W2aR+x61yoTdca43KSZtzXZ+\n" +
	"kQT3ZczXbC47fI4FfR6GAkgDAoGBALIKusCNGsEd21gMp/C36hmtFYlpDnWTaDNh\n" +
	"kZ0yafuy+fBRJGQmuq4K1p40t4RB94rVSUMcn/yongeiSnx0Kjfo/jvVl3hks1Fv\n" +
	"NAdjgseqxvNmiu/XR3h3j0opziU8EEP69MpOfTWQd8FsqCaNSSRr4bMFKarQBgD3\n" +
	"eDXPbafLAoGAeQtSbyYfqGWZ21B/UySbLRIxi5milRmBHJaAuRK1gKe1LWO4awmf\n" +
	"JVf4nwF+QRxO2lFlTH7z+VqaZWix8yaLGIOdGTuPolyu6GGH5+iwdo2dKUWT4stw\n" +
	"w1qfChAeHnXEO3+bUX+NmTiCLINSkiGss0/uFZWIj1ZQijSVxJk05Bg=\n" +
	"-----END RSA PRIVATE KEY-----"

// GetRSAPublicKey 获取公钥
func GetRSAPublicKey(publicKey *rsa.PublicKey) (string, error) {
	data, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		return "", err
	}
	block := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: data,
	}
	bytes := pem.EncodeToMemory(&block)
	return string(bytes), nil
}

// GetRSAPrivateKey 获取私钥
func GetRSAPrivateKey(privateKey *rsa.PrivateKey) string {
	data := x509.MarshalPKCS1PrivateKey(privateKey)
	block := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: data,
	}
	bytes := pem.EncodeToMemory(&block)
	return string(bytes)
}

// getPrivateKey 获取私钥对象，转换后的
func getPrivateKey() (*rsa.PrivateKey, error) {
	block, _ := pem.Decode([]byte(privateKeyPem))
	if block == nil {
		return nil, errors.New("private key error")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

// getPublicKey 获取公钥对象
func getPublicKey() (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(publicKeyPem))
	if block == nil {
		return nil, errors.New("public key error")
	}
	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	key := publicKey.(*rsa.PublicKey)
	return key, nil
}

// RSADecode /**
func RSADecode(cipher string) (*string, error) {
	bytes, err := base64.StdEncoding.DecodeString(cipher)
	if err != nil {
		return nil, err
	}
	privateKey, err := getPrivateKey()
	if err != nil {
		return nil, err
	}
	temp, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, bytes)
	if err != nil {
		return nil, err
	}
	res := string(temp)
	return &res, err
}

// RsaEncode /**
func RsaEncode(plain string) (*string, error) {
	publicKey, err := getPublicKey()
	if err != nil {
		return nil, err
	}
	temp, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(plain))
	if err != nil {
		return nil, err
	}
	res := base64.StdEncoding.EncodeToString(temp)
	return &res, err
}
