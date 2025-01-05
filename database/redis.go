package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

// InitRedis, Redis bağlantısını başlatır
func InitRedis() error {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	// Redis bağlantı adresi oluştur
	addr := fmt.Sprintf("%s:%s", redisHost, redisPort)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisPassword,
		DB:       0, // varsayılan DB
	})

	// Bağlantıyı test et
	if err := RedisClient.Ping(context.Background()).Err(); err != nil {
		log.Printf("Redis bağlantı hatası: %v", err)
		return err
	}

	log.Println("Redis bağlantısı başarıyla kuruldu")
	return nil
}
