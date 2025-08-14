package builder

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizeJSON(t *testing.T) {
	// Sample JSON input (as bytes) that simulates a realistic nested structure
	// with organizations containing areas.
	jsonStr := []byte(`{
		"organizations": [
			{
			"orgID": 1,
			"orgName": "helsinki",
			"orgNameFi": "Helsinki",
			"orgNameSe": "Helsingfors",
			"areas": [
				{
				"areaID": 1,
				"areaNameFi": "Helsinki",
				"areaNameSe": "Helsingfors"
				}
			]
			}
		]
	}`)

	// NormalizeJSON parses and normalizes jsonStr
	v, err := NormalizeJSON(jsonStr)
	if err != nil {
		t.Fatalf("NormalizeJSON error: %v", err)
	}

	// Assert key is string, value is any
	root, ok := v.(map[string]any)
	if !ok {
		t.Fatalf("expected map[string]any at root, got %T", v)
	}

	assert.IsType(t, []any{}, root["organizations"], "organizations should be an array")
	assert.NotEmpty(t, root["organizations"].([]any), "organizations array should not be empty")


	// Take the first organization object from the array
	orgs := root["organizations"].([]any) // type assert to slice
	first, ok := orgs[0].(map[string]any)     // type assert the first element
	if !ok {
		t.Fatalf("expected organizations[0] to be object, got %T", orgs[0])
	}

	assert.Equal(t, float64(1), first["orgID"], "orgID should be `1`")
	assert.Equal(t, "helsinki" , first["orgName"] , "orgName should be `helsinki`")
	assert.Equal(t, "Helsingfors" , first["orgNameSe"] , "orgNameSe should be `Helsingfors`")

	fmt.Println(v)
}
