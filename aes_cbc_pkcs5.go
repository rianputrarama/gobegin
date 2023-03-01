package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

/**
 * Created by GoLand
 * Author : Febriansyah Putra Ramadhan
 * Date : 26/01/2023;
 * Time : 3:02;
**/

var (
	initialVector = "1234567890123456"
	passphrase    = "rctYFueGKyCRmkGFJd82LybBTXUMZoJX"
)

func main() {
	var plainText = "tester12_"

	encryptedData := AESEncrypt(plainText, []byte(passphrase))
	encryptedString := base64.StdEncoding.EncodeToString(encryptedData)
	fmt.Println(encryptedString)

	encryptedData, _ = base64.StdEncoding.DecodeString(encryptedString)
	decryptedText := AESDecrypt(encryptedData, []byte("0V4CzDUX207ACHYLH6lUywQdWY1IwVmdLA=="))
	fmt.Println(string(decryptedText))
}

func AESEncrypt(src string, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(initialVector))
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	return crypted
}

func AESDecrypt(crypt []byte, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if len(crypt) == 0 {
		fmt.Println("plain content empty")
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(initialVector))
	decrypted := make([]byte, len(crypt))
	ecb.CryptBlocks(decrypted, crypt)

	return PKCS5Trimming(decrypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
