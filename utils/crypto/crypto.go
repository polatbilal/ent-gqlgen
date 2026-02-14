package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Encrypt encrypts plaintext using AES-256-GCM
// Returns base64 encoded ciphertext
func Encrypt(plaintext string, key []byte) (string, error) {
	if len(plaintext) == 0 {
		return "", nil
	}

	// AES-256 için key 32 byte olmalı
	if len(key) != 32 {
		return "", errors.New("encryption key must be exactly 32 bytes for AES-256")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Nonce oluştur (GCM için)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Şifrele
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Base64 encode et
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt decrypts base64 encoded ciphertext using AES-256-GCM
// Returns plaintext
func Decrypt(ciphertext string, key []byte) (string, error) {
	if len(ciphertext) == 0 {
		return "", nil
	}

	// AES-256 için key 32 byte olmalı
	if len(key) != 32 {
		return "", errors.New("decryption key must be exactly 32 bytes for AES-256")
	}

	// Base64 decode
	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return "", errors.New("ciphertext too short")
	}

	nonce, cipherData := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
