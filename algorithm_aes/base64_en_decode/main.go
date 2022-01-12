package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

/*
对称加密：
	公钥加密，私钥解密，公钥无法解密
	私钥加密，公钥解密，私钥无法解密
*/

// 加密 aes_128_cbc
func Encrypt(encryptStr string, key []byte, iv string) (string, error) {
	encryptBytes := []byte(encryptStr)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockSize := block.BlockSize()
	encryptBytes = pkcs5Padding(encryptBytes, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(encryptBytes))
	blockMode.CryptBlocks(encrypted, encryptBytes)
	return base64.URLEncoding.EncodeToString(encrypted), nil
}

// 解密
func Decrypt(decryptStr string, key, iv []byte) (string, error) {
	decryptBytes, err := base64.StdEncoding.DecodeString(decryptStr)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, iv)
	decrypted := make([]byte, len(decryptBytes))

	blockMode.CryptBlocks(decrypted, decryptBytes)
	decrypted = pkcs5UnPadding(decrypted)
	return string(decrypted), nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcs5UnPadding(decrypted []byte) []byte {
	length := len(decrypted)
	unPadding := int(decrypted[length-1])
	return decrypted[:(length - unPadding)]
}

func main() {
	originData := "qnZs44S6SmkRy9LZr8I7o1ZSbbEMbQekOr4ANTUUiJEsi6hBXwu2ORmFz5UKYSkC5ROyORoBEv7OHUfMITczAOcrOdK62gFeh06tYOaBU4/jEt+12fow3pro/eFDLRwVI8Mcc7pNnhjR4EtjK6HRMQ=="
	keyBytes, _ := base64.StdEncoding.DecodeString("M5QTNzMjN4IjNwYDO54CM3cTNz4SNzUTO5ETOxITM1g=")
	decryptData, err := Decrypt(originData, keyBytes, keyBytes[:16])
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("decrypt data: %v", decryptData)

}
