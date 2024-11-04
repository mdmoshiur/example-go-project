package dto

import "encoding/json"

// Object represents a common type
type Object map[string]interface{}

// Encode encodes a struct to object
func (o *Object) Encode(i interface{}) error {
	jsonByte, err := json.Marshal(i)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonByte, o)

	return err
}

// Decode decodes a object to struct
func (o *Object) Decode(i interface{}) error {
	jsonByte, err := json.Marshal(o)
	if err != nil {
		return err
	}

	err = json.Unmarshal(jsonByte, i)

	return err
}

// String converts object to json string
func (o *Object) String() (string, error) {
	jsonByte, err := json.Marshal(o)
	if err != nil {
		return "", nil
	}
	return string(jsonByte), nil
}
