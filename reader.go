// Laraboot shared tools
package shared

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

// NewFromString : Reads JSON data from string and returns the parsed map.
func NewFromString(jsonData string) (
	LarabootStruct,
	error,
) {
	var m LarabootStruct
	reader := strings.NewReader(jsonData)

	dec := json.NewDecoder(reader)

	if err := dec.Decode(&m); err == io.EOF {
		fmt.Printf("unmarshallErr: %v", err)
		return m, err
	}

	return m, nil
}

// NewFromFile :Reads JSON data from file and returns the parsed map.
func NewFromFile(filename string) (LarabootStruct, error) {
	var m LarabootStruct
	file, errReadingJSONFile := os.Open(filename) //nolint:gosec //cus

	if errReadingJSONFile != nil {
		return m, errReadingJSONFile
	}

	errDecoding := json.NewDecoder(file).Decode(&m)

	if errDecoding != nil {
		fmt.Printf("unmarshallErr: %v", errDecoding)
	}

	return m, nil
}
