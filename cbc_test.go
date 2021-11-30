package aes

import (
	"fmt"
	"testing"
)

func TestAesEncryptCBC(t *testing.T) {
	orig := "hello world"
	key := "0123456789012345"
	fmt.Println("origin：", orig)
	encryptCode := AesEncryptCBC(orig, key)
	fmt.Println("encrypt：", encryptCode)
	decryptCode := AesDecryptCBC(encryptCode, key)
	fmt.Println("decrypt：", decryptCode)
}
