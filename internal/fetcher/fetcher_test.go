package fetcher

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestFetchJSON(t *testing.T) {
	url := os.Getenv("URL")
	if strings.Compare(url, "") == 0 {
		log.Fatal("URL is not set up")
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
