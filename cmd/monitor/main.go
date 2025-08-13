package main

import (
	"context"
	"fmt"
	// "fmt"
	// "io"
	"os"
	"log"
	// "net/http"

	// "github.com/redis/go-redis/v9"
	"github.com/joho/godotenv"

	"github.com/iliamunaev/monitor-wep-page/internal/fetcher"
)

var ctx = context.Background()

func main() {

	// Fetch JSON
	_ = godotenv.Load("../../.env")
	url := os.Getenv("SERVICE_URL")

    data, err := fetcher.FetchJSON(url)
    if err != nil {
        log.Fatal("Failed to fetch JSON:", err)
    }
    fmt.Println(string(data))

    // Upload to Redis

    // Compare

    // Notify if changes

}
