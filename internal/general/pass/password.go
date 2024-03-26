package pass

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"io"
)

type HashData struct {
	SecretKey string
	Data      string
	NonceHex  string
}

func HashH512Password(has HashData) string {
	pwdByte := []byte(has.Data)

	salt := []byte(has.SecretKey)

	sha512 := sha512.New()

	pwdByte = append(pwdByte, salt...)

	sha512.Write(pwdByte)

	hashedPassword := sha512.Sum(nil)

	return hex.EncodeToString(hashedPassword)
}

func ComparePasswordAndHash(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func ComparePasswordHash(passwordHashRPC, passwordHashDB string) bool {
	return passwordHashRPC == passwordHashDB
}

func Encrypt(data HashData) (*HashData, error) {
	normalizeKey := sha256.Sum256([]byte(data.SecretKey))

	block, err := aes.NewCipher(normalizeKey[:])
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceBytes := make([]byte, aesgcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonceBytes); err != nil {
		return nil, err
	}

	encryptedBytes := aesgcm.Seal(nil, nonceBytes, []byte(data.Data), nil)
	return &HashData{Data: hex.EncodeToString(encryptedBytes), NonceHex: hex.EncodeToString(nonceBytes)}, nil
}

func Decrypt(data HashData) (*HashData, error) {
	normalizeKey := sha256.Sum256([]byte(data.SecretKey))

	block, err := aes.NewCipher(normalizeKey[:])
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonceBytes, err := hex.DecodeString(data.NonceHex)
	if err != nil {
		return nil, err
	}

	encryptedBytes, err := hex.DecodeString(data.Data)
	if err != nil {
		return nil, err
	}

	decryptedBytes, err := aesgcm.Open(nil, nonceBytes, encryptedBytes, nil)
	if err != nil {
		return nil, err
	}

	return &HashData{Data: string(decryptedBytes)}, nil
}
