package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func mdHashing(input string) string {
	byteInput := []byte(input)
	md5Hash := md5.Sum(byteInput)
	return hex.EncodeToString(md5Hash[:]) // by referring to it as a string
}

func AESEncrypt(value []byte, keyPhrase string) ([]byte, error) {
	aesBlock, err := aes.NewCipher([]byte(mdHashing(keyPhrase)))
	if err != nil {
		return []byte{}, err
	}

	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return []byte{}, err
	}

	nonce := make([]byte, gcmInstance.NonceSize())
	_, _ = io.ReadFull(rand.Reader, nonce)

	cipheredText := gcmInstance.Seal(nonce, nonce, value, nil)

	return cipheredText, nil
}

func AESDecryptIt(ciphered []byte, keyPhrase string) ([]byte, error) {
	hashedPhrase := mdHashing(keyPhrase)
	aesBlock, err := aes.NewCipher([]byte(hashedPhrase))
	if err != nil {
		return []byte{}, err
	}
	gcmInstance, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return []byte{}, err
	}

	nonceSize := gcmInstance.NonceSize()
	nonce, cipheredText := ciphered[:nonceSize], ciphered[nonceSize:]

	originalText, err := gcmInstance.Open(nil, nonce, cipheredText, nil)
	if err != nil {
		return []byte{}, err
	}

	return originalText, nil
}
