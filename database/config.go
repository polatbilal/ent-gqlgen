package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/polatbilal/ent-gqlgen/ent"

	// _ "github.com/go-sql-driver/mysql" //MySQL driver
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
	// databaseURL := os.Getenv("MYSQL_URL")
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Printf("Error: DATABASE_URL environment variable is not set")
		return nil, fmt.Errorf("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", databaseURL)
	//db, err := sql.Open("mysql", databaseURL)
	if err != nil {
		log.Printf("Error opening database: %v\n", err)
		return nil, err
	}

	// Bağlantı havuzu yapılandırması - server ortamı için optimize edildi
	db.SetMaxIdleConns(10)                  // Batch işlemler için yeterli boşta bağlantı
	db.SetMaxOpenConns(50)                  // Eşzamanlı batch işlemler için artırıldı
	db.SetConnMaxLifetime(5 * time.Minute)  // Daha uzun yaşam süresi
	db.SetConnMaxIdleTime(90 * time.Second) // Transaction'lar için yeterli süre

	// Bağlantı havuzunu test et
	if err = db.Ping(); err != nil {
		log.Printf("Error pinging database: %v\n", err)
		return nil, err
	}

	// Ent istemcisini yapılandırılmış SQL bağlantısı ile oluştur
	drv := entsql.OpenDB("postgres", db)
	// drv := entsql.OpenDB("mysql", db)
	client := ent.NewClient(ent.Driver(drv))

	return client, nil
}
