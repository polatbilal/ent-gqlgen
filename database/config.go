package database

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/polatbilal/gqlgen-ent/ent"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var (
	dbClients = make(map[string]*ent.Client)
	mu        sync.Mutex
)

func GetClient(companyCode string) (*ent.Client, error) {
	mu.Lock()
	defer mu.Unlock()

	if client, exists := dbClients[companyCode]; exists {
		return client, nil
	}

	client, err := Connect(companyCode)
	if err != nil {
		return nil, err
	}

	dbClients[companyCode] = client
	return client, nil
}

func Connect(companyCode string) (*ent.Client, error) {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Printf("Error: DATABASE_URL environment variable is not set")
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	client, err := ent.Open("postgres", databaseURL)
	if err != nil {
		log.Printf("Error: postgres client: %v\n", err)
		return nil, err
	}

	return client, nil
}
