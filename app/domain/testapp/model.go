package testapp

import "encoding/json"

type status struct {
	Status string
}

func (s status) Encode() ([]byte, string, error) {
	data, err := json.Marshal(s)
	return data, "application/json", err
}
