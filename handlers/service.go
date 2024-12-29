package handlers

import "net/http"

type ExternalService struct {
	baseURL string
	client  *http.Client
}
