package crypto

import (
	"fmt"
	"github.com/marspere/goencrypt"
)
var PassWord  = "mZKM454z"
func DesEncode(str string) (string, error) {
	// key为12345678
	// iv为空
	// 采用ECB分组模式
	// 采用pkcs5padding填充模式
	// 输出结果使用base64进行加密
	cipher := goencrypt.NewDESCipher([]byte(PassWord), []byte(PassWord), goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintHex)
	cipherText, err := cipher.DESEncrypt([]byte(str))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return cipherText, nil
}
func DesDecode(str string) (string, error) {
	// key为12345678
	// iv为空
	// 采用ECB分组模式
	// 采用pkcs5padding填充模式
	// 输出结果使用base64进行加密
	cipher := goencrypt.NewDESCipher([]byte(PassWord), []byte(PassWord), goencrypt.CBCMode, goencrypt.Pkcs7, goencrypt.PrintHex)
	cipherText, err := cipher.DESDecrypt(str)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return cipherText, nil
}
