package helper

import (
	"encoding/json"
	"log"
)

// MarshalJSON is a generic function to marshal any data structure to JSON
func MarshalJSON(data interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %s", err)
		return nil, err
	}
	return jsonData, nil
}
