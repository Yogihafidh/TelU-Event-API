package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "secret-key"

func GenerateToken(email string, userId int64) (string, error) {
	// Create a new token object. jwt.NewWithClaims akan membuat token baru dengan klaim yang diberikan.
	// Klaim adalah informasi yang disimpan dalam token, seperti email, userId, dan waktu kedaluwarsa.
	// Dalam hal ini, kita menggunakan jwt.MapClaims untuk menyimpan klaim dalam bentuk map
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(), // Token berlaku selama 24 jam
	})

	// convert token to string using SignedString method. SignedString akan mengembalikan string token yang sudah ditandatangani dengan secretKey
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (string, int64, error) {
	// 1. Parse token and verify signature. jwt.Parse akan mengembalikan token yang sudah diverifikasi
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		// Check if the signing method is HMAC (HS256)
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	// 2. If there is an error while parsing the token or expired, return an error
	if err != nil {
		return "", 0, errors.New("Could not parse token")
	}

	// 3. Check if the token is valid. Validasi token dengan memeriksa apakah token sudah kedaluwarsa atau tidak
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", 0, errors.New("Token is not valid")
	}

	// 4. Get data from the token. Klaim adalah informasi yang disimpan dalam token, seperti email, userId, dan waktu kedaluwarsa
	// jwt.MapClaims adalah tipe data yang digunakan untuk menyimpan klaim dalam bentuk map
	// In this case, we are extracting the email and userId from the claims
	// We also check if the claims can be converted to jwt.MapClaims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", 0, errors.New("Invalid token claims")
	}
	email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return email, userId, nil
}
