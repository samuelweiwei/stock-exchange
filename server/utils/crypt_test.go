package utils

import (
	"fmt"
	"testing"
)

func TestCrypt(t *testing.T) {
	key := []byte("2TPN4fApRAeCzWCy") // 16字节的密钥
	plaintext := []byte("3")

	// 加密
	ciphertext, err := Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("Encrypt error:", err)
		return
	}
	fmt.Println("Ciphertext:", ciphertext)

	// 解密
	decryptBytes, err := Decrypt([]byte(ciphertext), key)
	if err != nil {
		fmt.Println("Decrypt error:", err)
		return
	}
	fmt.Println("Decrypted:", string(decryptBytes))

}
