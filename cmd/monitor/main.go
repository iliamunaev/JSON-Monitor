package main

import (
	"log"
	"os"
	"reflect"
	"time"

	"github.com/iliamunaev/json-monitor/internal/builder"
	"github.com/iliamunaev/json-monitor/internal/fetcher"
	"github.com/joho/godotenv"
)

func init() {
	_ = godotenv.Load("../../.env")
}

func main() {
	url, ok := os.LookupEnv("SERVICE_URL")
	if !ok || url == "" {
		log.Fatal("Error: SERVICE_URL is not set")
	}

	interval := 300 * time.Second

	var prev any

	for {
		raw, err := fetcher.FetchJSON(url)
		if err != nil {
			log.Printf("Fetch error: %v", err)
			time.Sleep(interval)
			continue
		}

		cur, err := builder.NormalizeJSON(raw)
		if err != nil {
			log.Printf("Normalize error: %v", err)
			time.Sleep(interval)
			continue
		}

		if prev == nil {
			log.Println("Initial fetch complete")
		} else if !reflect.DeepEqual(prev, cur) {
			log.Println("Changes detected!")

		} else {
			log.Println("No changes")
		}

		prev = cur
		time.Sleep(interval)
	}
}
