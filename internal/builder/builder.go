package builder

import (
	"encoding/json"
	"fmt"
)

func NormalizeJSON(b []byte) (any, error) {
	var v any
	err := json.Unmarshal(b, &v)
	if err != nil {
		fmt.Println("Error: normalize json:", err)
		return v, err
	}
	return v, nil
}
