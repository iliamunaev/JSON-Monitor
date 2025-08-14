package comparator

import (
	"log"
	"reflect"
)

func CompareJSONs(j1, j2 any) (bool, error) {
	if j1 == nil && j2 == nil {
		log.Println("No changes (both nil)")
		return true, nil
	}

	if j1 == nil || j2 == nil {
		log.Println("Changes detected! (one is nil)")
		return false, nil
	}

	if reflect.DeepEqual(j1, j2) {
		log.Println("No changes")
		return true, nil
	} else {
		log.Println("Changes detected!")
		return false, nil
	}
}

