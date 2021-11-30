package aes

import (
	"encoding/hex"
	"fmt"
	"log"
	"testing"
)

func TestAesEncryptCFB(t *testing.T) {
	source := "hello world"
	fmt.Println("origin：", source)
	key := "ABCDEFGHIJKLMNO1" // 16 byte

	encryptCode, err := AesEncryptCFB([]byte(source), []byte(key))
	if err != nil {
		log.Printf("err:%v", err)
		return
	}

	fmt.Println("encryptCode：", hex.EncodeToString(encryptCode))
	decryptCode, err := AesDecryptCFB(encryptCode, []byte(key))
	if err != nil {
		log.Printf("err:%v", err)
		return
	}

	fmt.Println("decryptCode：", string(decryptCode))
}
