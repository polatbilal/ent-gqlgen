package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type GraphQLClient struct {
	URL string
}

type graphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// NewGraphQLClient creates a new GraphQL client with the configured URL
func NewGraphQLClient(scheme string) *GraphQLClient {
	return &GraphQLClient{
		URL: os.Getenv("GRAPHQL_URL"),
	}
}

func (c *GraphQLClient) Execute(query string, variables map[string]interface{}, token string, result interface{}) error {
	requestBody := graphQLRequest{
		Query:     query,
		Variables: variables,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("request body oluşturma hatası: %v", err)
	}

	// log.Printf("GraphQL Request: %s", string(jsonBody))

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

	// log.Printf("GraphQL Response: %s", string(body))

	var response struct {
		Data   json.RawMessage `json:"data"`
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

	if err := json.Unmarshal(response.Data, result); err != nil {
		return fmt.Errorf("data parse hatası: %v", err)
	}

	return nil
}
