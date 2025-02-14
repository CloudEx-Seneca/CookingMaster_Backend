package authhelper

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strconv"
	"strings"
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
	var time, memory uint32
	var threads uint8
	var saltB64, hashB64 string

	_, time, memory, threads, saltB64, hashB64, err := parseEncodedHash(pm.encodedHash)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(saltB64)
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

func parseEncodedHash(encodedHash string) (string, uint32, uint32, uint8, string, string, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return "", 0, 0, 0, "", "", fmt.Errorf("invalid encoded hash format")
	}

	time, err := strconv.ParseUint(parts[1], 10, 32)
	if err != nil {
		return "", 0, 0, 0, "", "", fmt.Errorf("invalid time: %v", err)
	}
	memory, err := strconv.ParseUint(parts[2], 10, 32)
	if err != nil {
		return "", 0, 0, 0, "", "", fmt.Errorf("invalid memory: %v", err)
	}
	threads, err := strconv.ParseUint(parts[3], 10, 8)
	if err != nil {
		return "", 0, 0, 0, "", "", fmt.Errorf("invalid threads: %v", err)
	}

	return parts[0], uint32(time), uint32(memory), uint8(threads), parts[4], parts[5], nil
}
