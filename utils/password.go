package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

const (
	argonTime    = 1
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

	// fungsi Argon2id, salah satu algoritma hashing password teraman saat ini. Hasilnya Hash password dalam bentuk []byte.
	hash := argon2.IDKey([]byte(password), salt, argonTime, memory, threads, keyLen)

	// Mengubah salt dan hash ke string agar mudah disimpan di database karena database tidak bisa menyimpan byte array secara langsung
	saltB64 := base64.RawStdEncoding.EncodeToString(salt)
	hashB64 := base64.RawStdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s$%s", saltB64, hashB64), nil
}

// VerifyPassword membandingkan input password dengan hash yang sudah disimpan (encodedHash)
func VerifyPassword(password, encodedHash string) (bool, error) {
	// Pisahkan encodedHash menjadi dua bagian: salt dan hash, yang dipisahkan oleh simbol "$"
	parts := strings.Split(encodedHash, "$")
	
	// Validasi: encodedHash harus terdiri dari 2 bagian saja (salt dan hash)
	if len(parts) != 2 {
		return false, errors.New("invalid hash format") // format tidak valid
	}

	// Decode bagian pertama (salt) dari base64 ke bentuk []byte
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false, err // gagal decode salt
	}

	// Decode bagian kedua (hash) dari base64 ke bentuk []byte
	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false, err // gagal decode hash
	}

	// Hash ulang password input dengan salt yang sama dan parameter yang sama (time, memory, threads, keyLen)
	hash := argon2.IDKey([]byte(password), salt, argonTime, memory, threads, keyLen)

	// Bandingkan hasil hash yang baru dengan hash yang disimpan menggunakan perbandingan waktu konstan (untuk mencegah timing attack)
	if subtle.ConstantTimeCompare(hash, expectedHash) == 1 {
		return true, nil // password cocok
	}

	// Jika tidak cocok, return false
	return false, nil
}