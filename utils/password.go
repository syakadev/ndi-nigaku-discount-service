package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"strings"

	"golang.org/x/crypto/argon2"
)

func GenerateSalt(size int) (string, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return "", errors.New("gagal menghasilkan salt")
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}

func HashPassword(password, username, salt string) (string, error) {
	saltBytes, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return "", errors.New("gagal mendekode salt")
	}

	combinedPassword := []byte(password + username)

	hash := argon2.IDKey(combinedPassword, saltBytes, 3, 64*1024, 4, 32)
	hashedBase64 := base64.StdEncoding.EncodeToString(hash)

	return hashedBase64, nil
}

func FormatHashedPassword(hash, salt string) string {
	return hash + ":" + salt
}
func VerifyPassword(password, username, storedHash string) (bool, error) {
	parts := strings.Split(storedHash, ":")
	if len(parts) != 2 {
		return false, errors.New("format password hash tidak valid")
	}

	hashedPasswordBase64, saltBase64 := parts[0], parts[1]

	saltBytes, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		return false, errors.New("gagal mendekode salt")
	}

	combinedPassword := []byte(password + username)
	newHashBytes := argon2.IDKey(combinedPassword, saltBytes, 3, 64*1024, 4, 32)
	newHashBase64 := base64.StdEncoding.EncodeToString(newHashBytes)

	if subtle.ConstantTimeCompare([]byte(newHashBase64), []byte(hashedPasswordBase64)) == 1 {
		return true, nil
	}

	return false, nil
}
