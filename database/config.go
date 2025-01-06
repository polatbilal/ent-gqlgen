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
	client *ent.Client
	once   sync.Once
)

func GetClient() (*ent.Client, error) {
	var err error
	once.Do(func() {
		client, err = Connect()
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func Connect() (*ent.Client, error) {
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
