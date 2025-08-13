package builder

import (
	"encoding/json"
	"fmt"
	"sort"
)

func NormalizeJSON(b []byte) (any, error) {
	var v any
	err := json.Unmarshal(b, &v)
	if err != nil {
		fmt.Println("Error: normalize json:", err)
		return v, err
	}
	v = sortMaps(v)
	return v, nil
}

func sortMaps(v any) any {
	switch x := v.(type) {
	case map[string]any:
		keys := make([]string, 0, len(x))
		for k := range x {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		sorted := make(map[string]any, len(x))
		for _, k := range keys {
			sorted[k] = sortMaps(x[k]) // recurse
		}
		return sorted

	case []any:
		for i := range x {
			x[i] = sortMaps(x[i]) // recurse into array elements
		}
		return x

	default:
		return v
	}
}
