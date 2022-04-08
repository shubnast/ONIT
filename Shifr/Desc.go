package Shifr

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Desc1(con *gin.Context) {

	key := []byte("12345678")
	iv := []byte("43218765")
	data, err := con.GetRawData()
	if err != nil {
		fmt.Println(err)
	}
	result, err := DesCBCEncrypt(data, key, iv)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("DescEnd.data", result, 0777)
	if err != nil {
		fmt.Println(err)
	}
	b := hex.EncodeToString(result)
	con.JSON(http.StatusOK, gin.H{"str": b})
}

func DesCBCEncrypt(data, key, iv []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}

	data = pkcs5Padding(data, block.BlockSize())
	cryptText := make([]byte, len(data))

	blockMode := cipher.NewCBCEncrypter(block, iv)
	blockMode.CryptBlocks(cryptText, data)
	return cryptText, nil
}

func pkcs5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}
