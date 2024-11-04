package validation

import "strings"

// Errors maps a string key to a list of values.
// It is typically used for query parameters and form values.
// Unlike in the http.Header map, the keys in a Errors map
// are case-sensitive.
type Errors map[string]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns
// the empty string. To access multiple values, use the map
// directly.
func (e Errors) Get(key string) string {
	if e == nil {
		return ""
	}

	return e[key]
}

// Set sets the key to value. It replaces any existing values.
func (e Errors) Set(key, value string) {
	e[key] = value
}

// Del deletes the values associated with key.
func (e Errors) Del(key string) {
	delete(e, key)
}

// IsNil report whether the errors is empty
func (e Errors) IsNil() bool {
	return len(e) == 0
}

func (e Errors) Error() string {
	var errs []string
	for _, err := range e {
		errs = append(errs, err)
	}
	return strings.Join(errs, ", ")
}
