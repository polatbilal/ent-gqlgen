package tools

import (
	"fmt"
	"log"
	"os"
	"time"
)

func LogBatchError(yibfNo int, errorType, message string) {
	// logs klasörünü oluştur (yoksa)
	if err := os.MkdirAll("logs", 0755); err != nil {
		log.Printf("Log klasörü oluşturulamadı: %v", err)
		return
	}

	// Log dosyasını aç (varsa append, yoksa oluştur)
	logFile, err := os.OpenFile("logs/batch_errors.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Log dosyası açılamadı: %v", err)
		return
	}
	defer logFile.Close()

	// Log mesajını formatla ve yaz
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMessage := fmt.Sprintf("[%s] YibfNo: %d | Tip: %s | %s\n", timestamp, yibfNo, errorType, message)

	if _, err := logFile.WriteString(logMessage); err != nil {
		log.Printf("Log dosyasına yazılamadı: %v", err)
	}
}
