package aes

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestAesEncryptOFB(t *testing.T) {
	source := "hello world"
	fmt.Println("origin：", source)
	key := "1111111111111111" //16byte or 32byte

	encryptCode, err := AesEncryptOFB([]byte(source), []byte(key))
	if err != nil {
		log.Printf("err:%v", err)
		return
	}

	fmt.Println("encryptCode：", hex.EncodeToString(encryptCode))

	decryptCode, err := AesDecryptOFB(encryptCode, []byte(key))
	if err != nil {
		log.Printf("err:%v", err)
		return
	}

	fmt.Println("decryptCode：", string(decryptCode))
}
