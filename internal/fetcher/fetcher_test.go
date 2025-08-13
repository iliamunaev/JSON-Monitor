package fetcher

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Warning: could not load .env file:", err)
	}
}

func TestFetchJSON(t *testing.T) {
	url := os.Getenv("SERVICE_URL")
	if url == "" {
		t.Fatal("SERVICE_URL is not set")
	}

	data, err := FetchJSON(url)
	if err != nil {
		t.Fatalf("Failed to fetch JSON: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("Received empty JSON")
	}

	fmt.Println(string(data))
}
