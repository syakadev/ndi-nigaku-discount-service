package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	encryptionKey []byte
)

func init() {
	keyHex := os.Getenv("ENCRYPTION_KEY")
	if keyHex == "" {
		panic("ENCRYPTION_KEY environment variable not set")
	}

	key, err := hex.DecodeString(keyHex)
	if err != nil {
		panic("ENCRYPTION_KEY is not a valid hex string: " + err.Error())
	}

	if len(key) != 32 {
		panic("Decoded ENCRYPTION_KEY must be 32 bytes long for AES-256")
	}
	encryptionKey = key
}

// Encrypt mengenkripsi teks menggunakan AES-GCM.
func Encrypt(plaintext string) (string, error) {
	if plaintext == "" {
		return "", nil
	}

	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())

	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// Decrypt mendekripsi teks yang dienkripsi dengan AES-GCM.
func Decrypt(ciphertext string) (string, error) {
	if ciphertext == "" {
		return "", nil
	}

	data, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64: %w", err)
	}

	block, err := aes.NewCipher(encryptionKey)
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

	nonce, ciphertextBytes := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

func DecryptLevelID(levelID string) (string, string, error) {
	var cleanUsername, cleanLevelID string
	decryptedLevelID, err := Decrypt(levelID)
	if err != nil {
		return "", "", err

	} else {
		partedLevelID := strings.Split(decryptedLevelID, "_")
		cleanUsername = partedLevelID[0]
		cleanLevelID = partedLevelID[1]
	}

	return cleanUsername, cleanLevelID, nil
}

func EncryptLevelID(username, levelID string) (string, error) {
	toEncrypt := username + "_" + levelID
	encrypted, err := Encrypt(toEncrypt)
	if err != nil {
		return "", err
	}

	return encrypted, nil
}
