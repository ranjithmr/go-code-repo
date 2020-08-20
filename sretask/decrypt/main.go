package main

import (
	"crypto/des"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	_ "os"
)

func main() {

	hexchars := ""
	// decoding the key value using hex package
	key, err := hex.DecodeString("")
	if err != nil {
		log.Fatal(err)
	}

	// Decrypt function takes the encrypted string and a key value which is slice to byte and returns plain text
	plaintext := Decrypt(hexchars, key)
	//printing to standard output
	fmt.Println(plaintext)
}

/* Decrypt function takes hexvalue and key, returns the plain text
First it uses challange function to decode hexvalue and uses
desDecrypt function with the decoded hexvalue to decrpt the string */

func Decrypt(hexchars string, key []byte) string {
	// challange function takes the encrypted value, decodes it and returns message
	message, err := challenge(hexchars)
	if err != nil {
		log.Fatal(err)
	}
	// desDecrypt function takes the decoded message and key and decrypt the message to plaintext using DES algorithm
	plaintext, err := desDecrypt(message, key)
	if err != nil {
		log.Fatal(err)
	}

	return string(plaintext)
}

/* challange function takes hexadecimal value,
decodes hexvalue and returns slice of byte */

func challenge(hexchars string) ([]byte, error) {
	if len(hexchars)%3 != 0 {
		return nil, errors.New("Invalid Hex string")
	}

	var tmp string
	// removing the values from hexvalue which were added while encrypting and storing it to tmp
	for i := 0; i < len(hexchars); i += 3 {
		tmp += hexchars[i+1 : i+3]
	}

	// decoding the formatted hex value using hex package
	message, err := hex.DecodeString(tmp)
	if err != nil {
		return nil, err
	}
	return message, nil
}

/* desDecrypt takes decoded message which is returned from challange function and key, returns decrypted slice of byte and error if any
we are using DES method to decrypt the string */

func desDecrypt(message, key []byte) ([]byte, error) {
	keyblock, err := des.NewCipher(key) // NewCipher creates and returns a cipher block
	if err != nil {
		return nil, err
	}

	decrypted := make([]byte, len(message))
	tmp := decrypted

	// creating a block size
	bs := keyblock.BlockSize()
	if len(message)%bs != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}
	// DES key must be 8 bytes so we are creating blocksize and decrypting 8 bytes in each iteration
	// decrypting the message using Decrypt method from DES and keyblock
	// decrypted value is stored in tmp which is pointing to decrypted slice
	// after the end of first iteration we are reslcing message, tmp from 8 to end
	for len(message) > 0 {
		keyblock.Decrypt(tmp, message[:bs])
		message = message[bs:]
		tmp = tmp[bs:]
	}
	return decrypted, nil
}
