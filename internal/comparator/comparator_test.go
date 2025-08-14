package comparator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareJSONs(t *testing.T) {
	tests := []struct {
		name     string
		jsonA    []byte
		jsonB    []byte
		expected bool
	}{
		{
			name: "Equal JSONs",
			jsonA: []byte(`{
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
			}`),
			jsonB: []byte(`{
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
			}`),
			expected: true,
		},
		{
			name: "Different JSONs",
			jsonA: []byte(`{
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
			}`),
			jsonB: []byte(`{
				"organizations": [
					{
					"orgID": 2,
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
			}`),
			expected: false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			ok, err := CompareJSONs(tc.jsonA, tc.jsonB)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, ok)
		})
	}
}
