package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"

    "fetcher"
)

var ctx = context.Background()

func main() {

	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Redis address
		Password: "",               // No password by default
		DB:       0,                // Default DB
	})

	// Fetch current JSON
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Failed to fetch JSON:", err)
	}
	defer resp.Body.Close()



	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body:", err)
	}

	// Get last stored JSON from Redis
	lastData, err := rdb.Get(ctx, "services.json:last").Result()
	if err == redis.Nil {
		fmt.Println("First run — storing data.")
	} else if err != nil {
		log.Fatal("Redis error:", err)
	} else if lastData != string(body) {
		fmt.Println("⚠ JSON has changed!")
	} else {
		fmt.Println("✅ No changes detected.")
	}

	// Save new JSON
	if err := rdb.Set(ctx, "services.json:last", body, 0).Err(); err != nil {
		log.Fatal("Failed to store JSON:", err)
	}
}
