package main

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

func TestMainIntegration(t *testing.T) {
	// Create a fake JSON API
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"school": "42"}`))
	})
	server := httptest.NewServer(handler)
	defer server.Close()

	os.Setenv("URL", server.URL)
	os.Setenv("SECONDS_TIMEOUT", "1") // 1 second interval for quick test

	// Run main() in a goroutine so can stop it
	done := make(chan struct{})
	go func() {
		main() // Will run indefinitely unless exit
		close(done)
	}()

	time.Sleep(3 * time.Second)

	// Stop the program (simulate CTRL+C or end test)
	// In a real app, you'd make main listen for context cancel
	t.Log("Integration test completed (main ran with fake API)")
}
