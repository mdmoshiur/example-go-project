package helper

import (
	"cmp"
	"encoding/json"
	"net/http"
)

// InSlice is a generic function which checks slice contains a value or not
func InSlice[T cmp.Ordered](what T, where []T) bool {
	for _, v := range where {
		if v == what {
			return true
		}
	}
	return false
}

// SliceDiff returns the elements in `a` that aren't in `b`.
func SliceDiff(a, b []uint32) []uint32 {
	mb := make(map[uint32]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}

	var diff []uint32
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}

	return diff
}

func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, err
	}
	
	return v, nil
}
