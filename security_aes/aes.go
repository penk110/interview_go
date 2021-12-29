package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"errors"
	"fmt"
)

type Cipher interface {
	Encrypt(crypt []byte) ([]byte, error)
	Decrypt(crypt []byte) ([]byte, error)
	GetEncryptParam() map[string]string
}

type aesCipher struct {
	key         []byte
	iv          []byte
	block       cipher.Block
	enBlockMode cipher.BlockMode
	deBlockMode cipher.BlockMode
}

func NewCipher(algorithm string) (Cipher, error) {
	var (
		c   Cipher
		err error
	)
	switch algorithm {
	case "AES256":
		// 配置文件获取
		key := []byte("123456789012345678901234567890AA")
		iv := key[:16]
		c, err = newAESCipher(key, iv)
	default:
		return nil, errors.New("invalid cipher type: " + algorithm)
	}
	if err != nil {
		return nil, err
	}

	return c, nil
}

func newAESCipher(key, iv []byte) (*aesCipher, error) {
	var err error
	a := &aesCipher{
		key: key,
		iv:  iv,
	}
	if a.block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}
	// AES分组长度为128位 blockSize=16 单位字节
	// 初始向量的长度必须等于块block的长度16字节
	a.enBlockMode = cipher.NewCBCEncrypter(a.block, a.key[:a.block.BlockSize()])
	a.deBlockMode = cipher.NewCBCDecrypter(a.block, a.key[:a.block.BlockSize()])
	return a, nil

}

func (ac *aesCipher) PKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padText...)
}

func (ac *aesCipher) PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

func (ac *aesCipher) Encrypt(origData []byte) ([]byte, error) {
	blockSize := ac.block.BlockSize()
	origData = ac.PKCS5Padding(origData, blockSize)
	crypt := make([]byte, len(origData))
	ac.enBlockMode.CryptBlocks(crypt, origData)
	return crypt, nil
}

func (ac *aesCipher) Decrypt(crypt []byte) ([]byte, error) {
	origData := make([]byte, len(crypt))
	ac.deBlockMode.CryptBlocks(origData, crypt)
	origData = ac.PKCS5UnPadding(origData)
	return origData, nil
}

func (ac *aesCipher) GetEncryptParam() map[string]string {
	enParam := map[string]string{
		"private_key": string(ac.key),
		"iv":          string(ac.iv),
	}
	return enParam
}

func main() {
	var (
		paramByte []byte
		deByte    []byte
	)
	key := []byte("123456789012345678901234567890AA")
	iv := []byte("1234567890123456")
	aesCipher, err := newAESCipher(key, iv)
	if err != nil {
		panic(err)
	}
	apiParams := map[string]string{
		"method":       "POST",
		"schema":       "http",
		"api":          "",
		"query_string": "",
	}
	if paramByte, err = json.Marshal(apiParams); err != nil {
		panic(err)
	}
	paramByte, err = aesCipher.Encrypt(paramByte)
	if err != nil {
		panic(err)
	}
	fmt.Println("Encrypt: ", string(paramByte))
	deByte, err = aesCipher.Decrypt(paramByte)
	if err != nil {
		panic(err)
	}
	fmt.Println("Decrypt: ", string(deByte))

	//ULVpiyoNcLMONrIlp8C6yagFEd0EK0voF1qLblPYTO8cIpMrTsfEOWB6pS7m+mSzRtFw+oBEQSTEfrw/9IXC6zQkJ5uRMKNIPqtuCPA4CKo=
	//U6TF0X/nGTxv+jmnDX9rn67liIn77stvfTxO7fNaYtdJ1qkMN8vEUXAfbOvw8v1vNR5SzBT2rKQt9zTD/lbBnymsVQV62XSS8ugoB/RbO7U=
	//fmt.Println(base64.StdEncoding.EncodeToString(paramByte))

	//py := `hvwtO0tUm7y6lqaq71ENfWAj5bgCQXkScWNb+OY3aoZolJpbXUCDS7JYsRW8zhh4ngP5nYhbtN441zUeRQu3g9Rjo6tpAGzMdM1GghzFuw1v4AHkbrXLX7AZo0FpHWwINvZVnPGfIOCsu7f2D3rW5GT3reXsteiZN5NTiAWiNuiuRz6ncHSQMeRDtwlv6WdW9uUdtEuq9ijOCZq8BsdY32etGSvOIaVUREjx7Gg2r4nHvY1xbvkKAVcQLLKB/12qt6niui2OriJp8BtjMPWd03Wudl00CDJfer/47dfwla18QtpaXitGUDO01eKUynzGErf7D3S0JTpbTGpG6iBdBg==`
	//d, err := base64.StdEncoding.DecodeString(py)
	//fmt.Println(string(d), err)

	//encryptStr := "mp1dEDw/+fr2Xja4Khl0xUA/cx5l46ff9USUEaPj5+0="
	//encryptStr := `hvwtO0tUm7y6lqaq71ENfWAj5bgCQXkScWNb+OY3aoZolJpbXUCDS7JYsRW8zhh4ngP5nYhbtN441zUeRQu3g9Rjo6tpAGzMdM1GghzFuw1v4AHkbrXLX7AZo0FpHWwINvZVnPGfIOCsu7f2D3rW5GT3reXsteiZN5NTiAWiNuiuRz6ncHSQMeRDtwlv6WdW9uUdtEuq9ijOCZq8BsdY32etGSvOIaVUREjx7Gg2r4nHvY1xbvkKAVcQLLKB/12qt6niui2OriJp8BtjMPWd03Wudl00CDJfer/47dfwla18QtpaXitGUDO01eKUynzGErf7D3S0JTpbTGpG6iBdBg==`
	//encryptByte, err := base64.StdEncoding.DecodeString(encryptStr)
	//if err != nil {
	//	panic(err)
	//}
	//decrypt, err := aesCipher.Decrypt(encryptByte)
	//if err != nil {
	//	panic(err)
	//	return
	//}
	//
	//fmt.Println(decrypt)
	//encryptByte1, err := base64.StdEncoding.DecodeString(string(decrypt))
	//fmt.Println(string(encryptByte1), err)

	data := map[string]interface{}{
		"name":        "测试",
		"description": "测试",
	}
	if paramByte, err = json.Marshal(data); err != nil {
		panic(err)
	}
	paramByte, err = aesCipher.Encrypt(paramByte)
	if err != nil {
		panic(err)
	}
	paramByte, err = aesCipher.Decrypt(paramByte)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(paramByte))

	//var m interface{}
	//if err = json.Unmarshal([]byte("SW50ZXJuYWwgU2VydmVyIEVycm9y"), &m); err != nil {
	//	panic(err)
	//}
	//fmt.Println(m)
}

