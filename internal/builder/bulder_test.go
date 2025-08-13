package builder

import (
	"fmt"
	"testing"
)

func TestNormalizeJSON(t *testing.T) {
	b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	v, err := NormalizeJSON(b)
	if err != nil {
		t.Fatalf("NormalizeJSON error: %v", err)
	}

	m, ok := v.(map[string]any)
	if !ok {
		t.Fatalf("Expected map[string]any, got %T", v)
	}
	if !ok {
		t.Fatalf("Expected map[string]any, got %T", v)
	}

	if m["Name"] != "Bob" {
		t.Errorf("Expected Name to be Bob, got %v", m["Name"])
	}

	if m["Food"] != "Pickle" {
		t.Errorf("Expected Food to be Pickle, got %v", m["Food"])
	}

	fmt.Println(v)
}
