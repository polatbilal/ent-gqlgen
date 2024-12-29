package main

import (
	"fmt"
	"time"
)

func date() {
	// Milisaniyeyi saniyeye çevir
	timestamp := int64(1267999200000 / 1000)

	// Unix zaman damgasını tarihe çevir
	t := time.Unix(timestamp, 0).UTC()

	y := time.Unix(timestamp, 0).Local()
	location, err := time.LoadLocation("Europe/Istanbul")
	if err != nil {
		fmt.Println("Hata:", err)
		return
	}

	z := time.Unix(timestamp, 0).In(location)

	fmt.Println("Tarih (UTC):", t)
	fmt.Println("Tarih (Local):", y)
	fmt.Println("Tarih (UTC+3):", z)
}

func main() {
	date()
}
