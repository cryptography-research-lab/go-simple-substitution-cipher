package main

import (
	"fmt"
	simple_substitution_cipher "github.com/cryptography-research-lab/go-simple-substitution-cipher"
)

func main() {

	// 创建一个随机秘钥，注意解密的时候要使用相同的秘钥来解密
	key := simple_substitution_cipher.RandomKey()
	plaintext := "CC11001100"
	fmt.Println("明文是： " + plaintext)
	ciphertext, err := simple_substitution_cipher.Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("加密错误： " + err.Error())
		return
	}
	fmt.Println("加密后： " + ciphertext)

	decryptResult, err := simple_substitution_cipher.Decrypt(ciphertext, key)
	if err != nil {
		fmt.Println("解密错误： " + err.Error())
		return
	}
	fmt.Println("解密结果： " + decryptResult)
	//  Output:
	//  明文是： CC11001100
	//  加密后： KK11001100
	//  解密结果： CC11001100

}
