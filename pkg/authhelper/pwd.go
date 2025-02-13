package authhelper

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
)

type PasswordManager struct {
	password    string
	encodedHash string
}

func NewPasswordEncoder(password string) *PasswordManager {
	return &PasswordManager{
		password: password,
	}
}

func (pm *PasswordManager) EncodedHash() error {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return err
	}

	time := uint32(3)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLen := uint32(32)
	hash := argon2.IDKey([]byte(pm.password), salt, time, memory, threads, keyLen)
	pm.encodedHash = fmt.Sprintf("%s$%d$%d$%d$%s$%s",
		"argon2id",
		time,
		memory,
		threads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return nil
}

func (pm *PasswordManager) GetEncodedHash() string {
	return pm.encodedHash
}

func NewPasswordDecoder(password string, encodedHash string) *PasswordManager {
	return &PasswordManager{
		password:    password,
		encodedHash: encodedHash,
	}
}

func (pm *PasswordManager) VerifyPassword() (bool, error) {
	var algo string
	var time, memory uint32
	var threads uint8
	var saltB64, hashB64 string

	_, err := fmt.Scanf(pm.encodedHash, "%s$%d$%d$%d$%s$%s", &algo, &time, &memory, &threads, &saltB64, &hashB64)
	if err != nil {
		return false, err
	}

	salt, err := base64.StdEncoding.DecodeString(saltB64)
	if err != nil {
		return false, err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(hashB64)
	if err != nil {
		return false, err
	}

	newHash := argon2.IDKey([]byte(pm.password), salt, time, memory, threads, uint32(len(expectedHash)))

	return string(newHash) == string(expectedHash), nil
}
