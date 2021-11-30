package aes

import (
	"fmt"
	"testing"
)

func TestAesEncryptECB(t *testing.T) {
	source := "hello world"
	fmt.Println("origin：", source)

	//16 byte
	key := "1443flfsaWfdas"
	encryptCode := AesEncryptECB([]byte(source), []byte(key))
	fmt.Println("encryptCode：", string(encryptCode))

	decryptCode := AesDecryptECB(encryptCode, []byte(key))

	fmt.Println("decryptCode：", string(decryptCode))
}
