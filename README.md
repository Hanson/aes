# Aes

a golang library to use aes encrypt easier.

# Install

```
go get github.com/hanson/aes
```

# Document

## CBC
```
orig := "hello world"
key := "0123456789012345"

encryptCode := AesEncryptCBC(orig, key)
fmt.Println("encrypt：" , encryptCode)
decryptCode := AesDecryptCBC(encryptCode, key)
fmt.Println("decrypt：", decryptCode)
```

## CFB
```
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
```

## CTR
```
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
```

## ECB
```
source := "hello world"
fmt.Println("origin：", source)

//16 byte
key := "1443flfsaWfdas"
encryptCode := AesEncryptECB([]byte(source), []byte(key))
fmt.Println("encryptCode：", string(encryptCode))

decryptCode := AesDecryptECB(encryptCode, []byte(key))

fmt.Println("decryptCode：", string(decryptCode))
```

## OFB
```
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
```