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
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := fmt.Sprintf("%s_%s", os.Getenv("DB_NAME"), "3")
	sslmode := "disable"

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslmode)

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error: postgres client: %v\n", err)
		return nil, err
	}

	return client, nil
}
