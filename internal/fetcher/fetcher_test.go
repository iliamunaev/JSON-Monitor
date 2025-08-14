package fetcher

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchJSON(t *testing.T) {
	// Create a fake JSON API endpoint
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"school":"42"}`))
	}))
	defer ts.Close()

	data, err := FetchJSON(ts.URL)
	if err != nil {
		t.Fatalf("Failed to fetch JSON: %v", err)
	}
	if len(data) == 0 {
		t.Fatal("Received empty JSON")
	}

	expected := `{"school":"42"}`
	if string(data) != expected {
		t.Fatalf("Expected %q, got %q", expected, string(data))
	}
}
