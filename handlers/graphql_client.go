package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type GraphQLClient struct {
	URL string
}

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

func (c *GraphQLClient) Execute(query string, variables map[string]interface{}, token string) error {
	requestBody := graphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("request body oluşturma hatası: %v", err)
	}

	log.Printf("GraphQL Request: %s", string(jsonBody))

	req, err := http.NewRequest("POST", c.URL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return fmt.Errorf("request oluşturma hatası: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request gönderme hatası: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("response okuma hatası: %v", err)
	}

	log.Printf("GraphQL Response: %s", string(body))

	var response struct {
		Data   interface{} `json:"data"`
		Errors []struct {
			Message string `json:"message"`
		} `json:"errors,omitempty"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("response parse hatası: %v", err)
	}

	if len(response.Errors) > 0 {
		return fmt.Errorf("GraphQL hatası: %s", response.Errors[0].Message)
	}

	return nil
}
