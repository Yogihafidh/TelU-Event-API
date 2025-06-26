package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	time    = 1
	memory  = 64 * 1024
	threads = 4
	keyLen  = 32
	saltLen = 16
)

// HashPassword hashes a password using Argon2id and returns base64(salt)$base64(hash)
func HashPassword(password string) (string, error) {
	// salt adalah byte array sepanjang saltLen.
	salt := make([]byte, saltLen)

	// mengisi array dengan nilai acak kriptografis.
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// fungsi dari algoritma Argon2id, salah satu algoritma hashing password teraman saat ini. Hasilnya Hash password dalam bentuk []byte.
	hash := argon2.IDKey([]byte(password), salt, time, memory, threads, keyLen)

	// Mengubah salt dan hash ke string agar mudah disimpan di database karena database tidak bisa menyimpan byte array secara langsung
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)
	hashB64 := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s$%s", saltB64, hashB64), nil
}
