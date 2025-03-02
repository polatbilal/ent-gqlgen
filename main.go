package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/polatbilal/gqlgen-ent/api-core"
	"github.com/polatbilal/gqlgen-ent/handlers-module"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Graceful shutdown için sinyal yakalama
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup

	// API Core modülünü başlat
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := api.StartServer(ctx); err != nil {
			log.Printf("API Core hatası: %v\n", err)
		}
	}()

	// Handlers modülünü başlat
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := handlers.StartServer(ctx); err != nil {
			log.Printf("Handlers modülü hatası: %v\n", err)
		}
	}()

	// Shutdown sinyalini bekle
	<-sigChan
	log.Println("Kapatma sinyali alındı. Servisler kapatılıyor...")
	cancel()

	// Tüm servislerin kapanmasını bekle
	wg.Wait()
	log.Println("Tüm servisler başarıyla kapatıldı.")
}
