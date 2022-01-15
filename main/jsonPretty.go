package main

import (
	"bytes"
	"encoding/json"
)

//Function to manually indent the JSON data for easy human readablity
func PrettyString(str string) (string, error) {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, []byte(str), "", "  "); err != nil {
		return "", err
	}
	return prettyJSON.String(), nil
}
