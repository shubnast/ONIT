package Shifr

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Aes1(con *gin.Context) {
	//FileName := "AesStart.txt"
	// logrus.Println("Encryption Program v0.01")
	//text, err := ioutil.ReadFile(FileName)
	//Попытка прочитать файл, если не вышло, выдача ошибки
	text, err := con.GetRawData()
	if err != nil {
		fmt.Println(err)
	}

	//ключ
	key := []byte("passphrasewhichneedstobe32bytes!")

	//создаем новый aes шифр по нашему ключу
	c, err := aes.NewCipher(key)
	//Ловим ошибки
	if err != nil {
		fmt.Println(err)
	}

	// gcm или Galois/Counter Mode, модификация
	// - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	gcm, err := cipher.NewGCM(c)
	//Ловим ошибки при генерации gcm
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcm.NonceSize())
	//Создаем одноразовый номер с случайной последовательностью
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	// here we encrypt our text using the Seal function
	// Seal encrypts and authenticates plaintext, authenticates the
	// additional data and appends the result to dst, returning the updated
	// slice. The nonce must be NonceSize() bytes long and unique for all
	// time, for a given key.

	//fmt.Println(gcm.Seal(nonce, nonce, text, nil))
	err = ioutil.WriteFile("AesEnd.data", gcm.Seal(nonce, nonce, text, nil), 0777)
	// Ловим ошибки записи
	if err != nil {
		fmt.Println(err)
	}

	con.JSON(http.StatusOK, gin.H{
		"str": gcm.Seal(nonce, nonce, text, nil),
	})
}
