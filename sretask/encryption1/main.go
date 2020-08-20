package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	text := []byte("")
	key := []byte("")

	ciphertext, err := Encrypt(text, key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", ciphertext)

	plaintext, err := Decrypt(text, key)
	if err != nil {
		log.Fatal(err)
	}
	decoded, err := hex.DecodeString(string(plaintext))
	if err != nil {
		log.Fatal(err)
	}
	return string(decoded)
	plaintext, _ := hex.DecodeString("")
	fmt.Printf("%s\n", plaintext)
}

func Encrypt(text, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, text, nil), nil
}

func Decrypt(text []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(text) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, text := text[:nonceSize], text[nonceSize:]
	return gcm.Open(nil, nonce, text, nil)
}
