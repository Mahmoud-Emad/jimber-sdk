// // This file just for testing.
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
)

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func encryptIt(value []byte, keyPhrase string) []byte {
	aesBlock, err := aes.NewCipher([]byte(mdHashing(keyPhrase)))
	if err != nil {
		fmt.Println(err)
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		fmt.Println(err)
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)

	return cipheredText
}

func decryptIt(ciphered []byte, keyPhrase string) []byte {
	hashedPhrase := mdHashing(keyPhrase)
	aesBlock, err := aes.NewCipher([]byte(hashedPhrase))
	if err != nil {
		log.Fatalln(err)
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		log.Fatalln(err)
	}

	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		log.Fatalln(err)
	}
	return originalText
}

func main() {
	myKey := "Mahmoud"
	myValue := "This is just a value"

	enc := encryptIt([]byte(myValue), myKey)
	fmt.Println(string(enc))
	dec := decryptIt(enc, myKey)
	fmt.Println(string(dec))
}

// import (
// 	"sync"

// 	app "github.com/Mahmoud-Emad/envserver/app"
// )

// func main() {
// 	server := app.NewServer("0.0.0.0", "8080")

// 	var wg sync.WaitGroup
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()
// 		if err := server.Serve(); err != nil {
// 		}
// 	}()
// 	wg.Wait()
// }
