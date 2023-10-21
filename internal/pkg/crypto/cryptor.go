package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// Key describes key of Encryption
type Key []byte

// Keys slice of Key
type Keys []Key

// Cryptor encryptor and decriptor and implemented AES Encryption with IV
// based on article https://ru.stackoverflow.com/a/863358
type Cryptor struct {
	keys    Keys
	lastKey byte
}

// NewCryptor ctor
func NewCryptor(keys Keys) Cryptor {
	return Cryptor{
		keys:    keys,
		lastKey: byte(len(keys) - 1),
	}
}

// Encrypt encrypts the message
func (tm Cryptor) Encrypt(message []byte) ([]byte, error) {
	key := tm.keys[tm.lastKey]
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	b := message
	b = pkcS5Padding(b, aes.BlockSize)
	encMessage := make([]byte, len(b))
	iv := key[:aes.BlockSize]
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(encMessage, b)

	return append(encMessage, tm.lastKey), nil
}

// Decrypt decrypts the message
func (tm Cryptor) Decrypt(encMessage []byte) ([]byte, error) {
	keyIdx := encMessage[len(encMessage)-1:][0]
	encMessage = encMessage[:len(encMessage)-1]
	key := tm.keys[keyIdx]
	iv := key[:aes.BlockSize]
	block, err := aes.NewCipher(key)

	if err != nil {
		return nil, err
	}

	if len(encMessage) < aes.BlockSize {
		return nil, errors.New("encMessage to short")
	}

	decrypted := make([]byte, len(encMessage))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, encMessage)

	return pkcS5UnPadding(decrypted), nil
}

func pkcS5Padding(cipher []byte, blockSize int) []byte {
	padding := blockSize - len(cipher)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(cipher, padText...)
}

func pkcS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])

	return src[:(length - unPadding)]
}
