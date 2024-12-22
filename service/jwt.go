package service

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type JwtCustomClaim struct {
	ID          int    `json:"userID"`
	Username    string `json:"username"`
	CompanyCode string `json:"companyCode"`
	jwt.StandardClaims
}

func JwtGenerate(ctx context.Context, userID int, username string, companyCode string) (string, error) {
	claims := &JwtCustomClaim{
		ID:          userID,
		Username:    username,
		CompanyCode: companyCode,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func JwtValidate(ctx context.Context, tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, fmt.Errorf("invalid token signature")
		}
		verr, ok := err.(*jwt.ValidationError)
		if ok && verr.Errors == jwt.ValidationErrorExpired {
			return nil, fmt.Errorf("token expired")
		}
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(*JwtCustomClaim); ok && token.Valid {
		if claims.ID == 0 || claims.Username == "" || claims.CompanyCode == "" {
			return nil, fmt.Errorf("invalid token claims")
		}
		return token, nil
	}

	return nil, fmt.Errorf("invalid token")
}
