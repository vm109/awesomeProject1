package amazon

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplaceZerosWithFivesInAInteger(t *testing.T) {
	value := ReplaceZerosWithFivesInAInteger(53420)
	assert.Equal(t, int32(53425), value)
}

func TestReplaceZerosWithFivesInAInteger2(t *testing.T) {
	value := ReplaceZerosWithFivesInAInteger(53400120)
	assert.Equal(t, int32(53455125), value)
}

func TestReplaceZerosWithFivesInAInteger3(t *testing.T) {
	value := ReplaceZerosWithFivesInAInteger(3400120)
	assert.Equal(t, int32(3455125), value)
}
