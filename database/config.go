package database

import (
	"fmt"
	"log"
	"sync"

	"gqlgen-ent/ent"

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
	user := "casaos"
	password := "casaos"
	host := "192.168.1.33"
	port := 5432
	dbName := fmt.Sprintf("gqlgen_%s", companyCode)
	sslmode := "disable"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbName, sslmode)

	client, err := ent.Open("postgres", connStr)
	if err != nil {
		log.Printf("Error: postgres client: %v\n", err)
		return nil, err
	}

	return client, nil
}
