package aes

import (
	"fmt"
	"log"
	"testing"
)

func TestAesCryptCTR(t *testing.T) {
	source := "hello world"
	fmt.Println("origin：", source)

	key := "1443flfsaWfdasds"
	encryptCode, err := AesCryptCTR([]byte(source), []byte(key))
	if err != nil {
		log.Printf("err:%v", err)
		return
	}

	fmt.Println("encryptCode：", string(encryptCode))

	decryptCode, err := AesCryptCTR(encryptCode, []byte(key))
	if err != nil {
		log.Printf("err:%v", err)
		return
	}

	fmt.Println("decryptCode：", string(decryptCode))
}
