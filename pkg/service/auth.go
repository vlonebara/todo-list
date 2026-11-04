package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const authTokenLifetime = 8 * time.Hour

type jwtHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type authClaims struct {
	PasswordHash string `json:"password_hash"`
	IssuedAt     int64  `json:"iat"`
	ExpiresAt    int64  `json:"exp"`
}

// Authenticate проверяет пароль и возвращает подписанный JWT-токен.
func Authenticate(password string) (string, error) {
	expectedPassword := os.Getenv("TODO_PASSWORD")
	if expectedPassword == "" {
		return "", errors.New("Пароль не настроен")
	}
	if !hmac.Equal([]byte(password), []byte(expectedPassword)) {
		return "", errors.New("Неверный пароль")
	}

	secret, err := jwtSecret()
	if err != nil {
		return "", err
	}
	return createAuthToken(expectedPassword, secret, time.Now())
}

// ValidateToken проверяет подпись, срок действия и парольную метку JWT.
func ValidateToken(token string) bool {
	password := os.Getenv("TODO_PASSWORD")
	secret, err := jwtSecret()
	if password == "" || err != nil {
		return false
	}

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return false
	}

	headerData, err := base64.RawURLEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}
	var header jwtHeader
	if err := json.Unmarshal(headerData, &header); err != nil || header.Algorithm != "HS256" || header.Type != "JWT" {
		return false
	}

	signature, err := base64.RawURLEncoding.DecodeString(parts[2])
	if err != nil || !hmac.Equal(signature, signJWT(parts[0]+"."+parts[1], secret)) {
		return false
	}

	claimsData, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}
	var claims authClaims
	if err := json.Unmarshal(claimsData, &claims); err != nil || claims.ExpiresAt <= time.Now().Unix() {
		return false
	}
	return hmac.Equal([]byte(claims.PasswordHash), []byte(passwordChecksum(password, secret)))
}

// jwtSecret возвращает ключ подписи из переменной TODO_JWT_SECRET.
func jwtSecret() ([]byte, error) {
	secret := os.Getenv("TODO_JWT_SECRET")
	if len(secret) < 32 {
		return nil, errors.New("TODO_JWT_SECRET должен содержать не менее 32 байт")
	}
	return []byte(secret), nil
}

// createAuthToken формирует JWT с меткой пароля и временем истечения.
func createAuthToken(password string, secret []byte, now time.Time) (string, error) {
	header, err := json.Marshal(jwtHeader{Algorithm: "HS256", Type: "JWT"})
	if err != nil {
		return "", fmt.Errorf("encode JWT header: %w", err)
	}
	claims, err := json.Marshal(authClaims{
		PasswordHash: passwordChecksum(password, secret),
		IssuedAt:     now.Unix(),
		ExpiresAt:    now.Add(authTokenLifetime).Unix(),
	})
	if err != nil {
		return "", fmt.Errorf("encode JWT claims: %w", err)
	}

	encodedHeader := base64.RawURLEncoding.EncodeToString(header)
	encodedClaims := base64.RawURLEncoding.EncodeToString(claims)
	signingInput := encodedHeader + "." + encodedClaims
	signature := base64.RawURLEncoding.EncodeToString(signJWT(signingInput, secret))
	return signingInput + "." + signature, nil
}

// passwordChecksum создаёт подписанную метку для проверки пароля при валидации.
func passwordChecksum(password string, secret []byte) string {
	checksum := hmac.New(sha256.New, secret)
	_, _ = checksum.Write([]byte("password-checksum:"))
	_, _ = checksum.Write([]byte(password))
	return base64.RawURLEncoding.EncodeToString(checksum.Sum(nil))
}

// signJWT подписывает данные алгоритмом HS256.
func signJWT(value string, secret []byte) []byte {
	signature := hmac.New(sha256.New, secret)
	_, _ = signature.Write([]byte(value))
	return signature.Sum(nil)
}
