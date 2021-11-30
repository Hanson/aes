package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func AesEncryptCBC(orig string, key string) string {
	origData := []byte(orig)
	k := []byte(key)

	block, _ := aes.NewCipher(k)

	blockSize := block.BlockSize()

	origData = PKCS7Padding(origData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])

	cryted := make([]byte, len(origData))

	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}

func AesDecryptCBC(cryted string, key string) string {
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)

	block, _ := aes.NewCipher(k)

	blockSize := block.BlockSize()

	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])

	orig := make([]byte, len(crytedByte))

	blockMode.CryptBlocks(orig, crytedByte)

	orig = PKCS7UnPadding(orig)
	return string(orig)
}
