package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

//加密
func AesEncryptSimple(origData []byte, key string, iv string) ([]byte, error) {
	return AesDecryptPkcs5(origData, []byte(key), []byte(iv))
}

func AesEncryptPkcs5(origData []byte, key []byte, iv []byte) ([]byte, error) {
	return AesEncrypt(origData, key, iv, PKCS5Padding)
}

func AesEncrypt(origData []byte, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = paddingFunc(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//解密
func AesDecryptSimple(crypted []byte, key string, iv string) ([]byte, error) {
	return AesDecryptPkcs5(crypted, []byte(key), []byte(iv))
}

func AesDecryptPkcs5(crypted []byte, key []byte, iv []byte) ([]byte, error) {
	return AesDecrypt(crypted, key, iv, PKCS5UnPadding)
}

func AesDecrypt(crypted, key []byte, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	fmt.Printf("%v\n", string(origData))
	origData = unPaddingFunc(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	if length < unpadding {
		return []byte("unpadding error")
	}
	return origData[:(length - unpadding)]
}

func main() {
	originStr := "qnZs44S6SmkRy9LZr8I7o1ZSbbEMbQekOr4ANTUUiJEsi6hBXwu2ORmFz5UKYSkC5ROyORoBEv7OHUfMITczAOcrOdK62gFeh06tYOaBU4/jEt+12fow3pro/eFDLRwVI8Mcc7pNnhjR4EtjK6HRMQ=="
	keyByte, _ := base64.StdEncoding.DecodeString("M5QTNzMjN4IjNwYDO54CM3cTNz4SNzUTO5ETOxITM1g=")
	originData, _ := base64.StdEncoding.DecodeString(originStr)
	decryptData, err := AesDecryptPkcs5(originData, keyByte, keyByte[:16])
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Printf("data: %v\n", string(decryptData))
}
