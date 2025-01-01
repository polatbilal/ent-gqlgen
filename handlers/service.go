package handlers

import "net/http"

type ExternalService struct {
	baseURL string
	client  *http.Client
}

type YDKTokenResponse struct {
	AccessToken    string         `json:"access_token"`
	TokenType      string         `json:"token_type"`
	RefreshToken   string         `json:"refresh_token"`
	DepartmentID   int            `json:"department_id"`
	ResponseSecret ResponseSecret `json:"response_secret"`
}

type ResponseSecret struct {
	Secret       string `json:"secret"`
	SecureSecret string `json:"secureSecret"`
	URI          string `json:"uri"`
}
