package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET environment variable is not set")
	}
	return []byte(secret)
}

type JwtCustomClaim struct {
	ID       int    `json:"userID"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func JwtGenerate(ctx context.Context, userID int, username string, name string, role string) (string, error) {
	claims := &JwtCustomClaim{
		ID:       userID,
		Username: username,
		Name:     name,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", fmt.Errorf("token imzalama hatası: %v", err)
	}

	return signedToken, nil
}

func JwtValidate(ctx context.Context, tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("beklenmeyen imzalama metodu: %v", t.Header["alg"])
		}
		return getJWTSecret(), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("geçersiz token imzası")
		}
		verr, ok := err.(*jwt.ValidationError)
		if ok && verr.Errors == jwt.ValidationErrorExpired {
			return nil, fmt.Errorf("token süresi dolmuş")
		}
		return nil, fmt.Errorf("geçersiz token")
	}

	if claims, ok := token.Claims.(*JwtCustomClaim); ok && token.Valid {
		if claims.ID == 0 || claims.Username == "" {
			return nil, fmt.Errorf("geçersiz token içeriği")
		}
		return token, nil
	}

	return nil, fmt.Errorf("geçersiz token")
}
