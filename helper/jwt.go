package helper

import (
	"fmt"
	"time"

	"os"

	"github.com/golang-jwt/jwt/v5"
)

// JWTClaims menyimpan data dalam token
type JWTClaims struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Username    string `json:"username"`
	NumberPhone string `json:"numberPhone"`
	Role        string `json:"role"`
	IssuedAt    int64  `json:"issued_at"`
	ExpiresAt   int64  `json:"expires_at"`
	jwt.RegisteredClaims
}

// GenerateJWT membuat token dengan data user dan masa berlaku 24 jam
func GenerateJWT(id, name, username, numberPhone, role string) (string, Error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", NewInternalServerError("jwt secret not found in environment")
	}

	expirationTime := time.Now().Add(24 * time.Hour).Unix()
	issueTime := time.Now().Unix()

	claims := JWTClaims{
		ID:          id,
		Name:        name,
		Username:    username,
		NumberPhone: numberPhone,
		Role:        role,
		IssuedAt:    issueTime,
		ExpiresAt:   expirationTime,
	}

	// Buat token dengan algoritma HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Tandatangani token dengan secret key
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", NewInternalServerError("failed to sign token")
	}

	return tokenString, nil
}

// ParseJWT memvalidasi token dan mengembalikan data user
func ParseJWT(tokenString string) (*JWTClaims, Error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return nil, NewInternalServerError("jwt secret not found in environment")
	}

	// Parse token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, NewUnauthenticatedError("Invalid token")
	}

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, NewUnauthenticatedError("Invalid token")
	}

	return claims, nil
}
