package simple_substitution_cipher

import (
	"github.com/golang-infrastructure/go-shuffle"
	"unicode/utf8"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// 检查key的合法性
func checkKey(key string) error {

	// 1. 长度必须为26位
	if len(key) != 26 {
		return ErrKeyUnavailable
	}

	// 2. 不允许有重复的映射，并且要恰好覆盖A-Z这26个字母
	// 先进行一次遍历统计
	charCount := make([]int, 26)
	for _, char := range key {
		// 统一转为小写字母
		char = toUppercaseIfNeed(char)
		if char < 'A' || char > 'Z' {
			return ErrKeyUnavailable
		}
		charCount[char-'A']++
	}

	// 然后再统一检查，A-Z是否每个都恰好出现一次
	for _, count := range charCount {
		if count != 1 {
			return ErrKeyUnavailable
		}
	}

	return nil
}

// 转换为大写字母，如果有必要的话
func toUppercaseIfNeed(char rune) rune {
	// 仅当为小写字母的时候才转换，其他字符的话就原样返回 
	if char >= 'a' && char <= 'z' {
		char -= 32
	}
	return char
}

// ------------------------------------------------ ---------------------------------------------------------------------

// RandomKey 生成一个随机的key
func RandomKey() string {

	// 先生成字母序列
	key := make([]rune, 26)
	for i := 0; i < 26; i++ {
		key[i] = 'A' + rune(i)
	}

	// 然后对序列进行shuffle
	shuffle.FisherYatesKnuthShuffle(key)

	return string(key)
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Encrypt 使用给定的key进行加密
func Encrypt(plaintext, key string) (string, error) {
	if err := checkKey(key); err != nil {
		return "", err
	}
	ciphertext := make([]rune, utf8.RuneCountInString(plaintext))
	for index, char := range plaintext {
		char = toUppercaseIfNeed(char)
		// 只有是字母的时候才进行加密
		if char >= 'A' && char <= 'Z' {
			ciphertext[index] = rune(key[char-'A'])
		} else {
			// 如果不是字母的话，则原样保留
			ciphertext[index] = char
		}
	}
	return string(ciphertext), nil
}

// ------------------------------------------------ ---------------------------------------------------------------------

// Decrypt 使用给定的key对密文进行解密
func Decrypt(ciphertext, key string) (string, error) {
	if err := checkKey(key); err != nil {
		return "", err
	}
	// 把key转为方便解密使用的形式
	key = convertKeyForDecrypt(key)
	plaintext := make([]rune, utf8.RuneCountInString(ciphertext))
	for index, char := range ciphertext {
		char = toUppercaseIfNeed(char)
		// 只有是字母的时候才进行加密
		if char >= 'A' && char <= 'Z' {
			plaintext[index] = rune(key[char-'A'])
		} else {
			// 如果不是字母的话，则原样保留
			plaintext[index] = char
		}
	}
	return string(plaintext), nil
}

// 把key转为方便解密使用的形式
func convertKeyForDecrypt(key string) string {
	decryptKey := make([]rune, 26)
	for index, char := range key {
		decryptKey[char-'A'] = rune('A' + index)
	}
	return string(decryptKey)
}

// ------------------------------------------------ ---------------------------------------------------------------------
