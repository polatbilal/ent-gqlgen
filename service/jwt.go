package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

var jwtSecret = []byte(getJwtSecret())
var tokenStore = NewTokenStore()

type JwtCustomClaim struct {
	ID          int    `json:"userID"`
	Username    string `json:"username"`
	CompanyCode string `json:"companyCode"`
	jwt.StandardClaims
}

// getJwtSecret .env dosyasından JWT_SECRET anahtarını yükler
func getJwtSecret() string {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "aSecret"
	}
	return secret
}

// JwtGenerate yeni bir token oluşturur ve varsa eski tokenları geçersiz kılar
func JwtGenerate(ctx context.Context, userID int, username string, companyCode string) (string, error) {
	claims := &JwtCustomClaim{
		ID:          userID,
		Username:    username,
		CompanyCode: companyCode,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(), // Token süresini 8 saat olarak ayarla
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	// Yeni token eklenirken eski tokenları geçersiz kıl
	tokenStore.AddToken(userID, signedToken)

	return signedToken, nil
}

// JwtValidate verilen tokenı doğrular ve geçerli olup olmadığını kontrol eder
func JwtValidate(ctx context.Context, tokenStr string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaim{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("invalid signing method")
		}
		return jwtSecret, nil
	})

	if err != nil {
		// Token geçersizse farklı hata döndürebilmek için hata mesajını kontrol edelim
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
		if !tokenStore.IsTokenValid(claims.ID, tokenStr) {
			return nil, fmt.Errorf("token invalidated")
		}
		return token, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// TokenStore kullanıcı tokenlarını saklar ve yönetir
type TokenStore struct {
	mu     sync.Mutex
	tokens map[int]string
}

func NewTokenStore() *TokenStore {
	return &TokenStore{
		tokens: make(map[int]string),
	}
}

func (store *TokenStore) AddToken(userID int, token string) {
	store.mu.Lock()
	defer store.mu.Unlock()
	store.tokens[userID] = token
}

func (store *TokenStore) IsTokenValid(userID int, token string) bool {
	store.mu.Lock()
	defer store.mu.Unlock()
	storedToken, exists := store.tokens[userID]
	return exists && storedToken == token
}

func (store *TokenStore) InvalidateToken(userID int) {
	store.mu.Lock()
	defer store.mu.Unlock()
	delete(store.tokens, userID)
}
