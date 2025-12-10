package services

import (
	"context"
	"errors"
	"time"

	"github.com/polatbilal/gqlgen-ent/ent"
	"github.com/polatbilal/gqlgen-ent/ent/user"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	client *ent.Client
	secret string
}

func NewAuthService(client *ent.Client, secret string) *AuthService {
	return &AuthService{
		client: client,
		secret: secret,
	}
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func (s *AuthService) Login(ctx context.Context, username, password string) (*TokenPair, error) {
	u, err := s.client.User.Query().
		Where(user.Username(username)).
		Only(ctx)
	if err != nil {
		return nil, errors.New("kullanıcı bulunamadı")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return nil, errors.New("geçersiz şifre")
	}

	// Access Token oluştur
	accessToken, err := s.generateAccessToken(u.ID)
	if err != nil {
		return nil, err
	}

	// Refresh Token oluştur
	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Refresh token'ı veritabanına kaydet
	expireAt := time.Now().Add(24 * time.Hour * 7) // 7 gün
	err = s.client.User.UpdateOne(u).
		SetRefreshToken(refreshToken).
		SetRefreshTokenExpireAt(expireAt).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error) {
	u, err := s.client.User.Query().
		Where(user.RefreshToken(refreshToken)).
		Only(ctx)
	if err != nil {
		return nil, errors.New("geçersiz refresh token")
	}

	if u.RefreshTokenExpireAt.Before(time.Now()) {
		return nil, errors.New("refresh token süresi dolmuş")
	}

	// Yeni token çifti oluştur
	accessToken, err := s.generateAccessToken(u.ID)
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	// Yeni refresh token'ı veritabanına kaydet
	expireAt := time.Now().Add(24 * time.Hour * 7) // 7 gün
	err = s.client.User.UpdateOne(u).
		SetRefreshToken(newRefreshToken).
		SetRefreshTokenExpireAt(expireAt).
		Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *AuthService) Logout(ctx context.Context, userID int) error {
	return s.client.User.UpdateOneID(userID).
		ClearRefreshToken().
		ClearRefreshTokenExpireAt().
		Exec(ctx)
}

func (s *AuthService) generateAccessToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.secret))
}

func (s *AuthService) generateRefreshToken() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(24 * time.Hour * 7).Unix() // 7 gün
	return token.SignedString([]byte(s.secret))
}
