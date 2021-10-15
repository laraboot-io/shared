package shared

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_larabootstruct_from_string(t *testing.T) { //nolint:unparam
	got, _ := NewFromString(`
    {"Name": "Alice", "Age": 25}
`)
	assert.ObjectsAreEqual(map[string]interface{}{"Age": 25, "Name": "Alice"}, got)
}

func Test_larabootstruct_from_file(t *testing.T) {
	got, _ := NewFromFile("./assets/laraboot.json")

	assert.Equal(t, "0.0.1", got.Version)
}
