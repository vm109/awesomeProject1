package google

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateBinaryStringFromGivenStringFirstTest(t *testing.T) {
	strings_filled := GenerateBinaryStringFromGivenString("1??0?101")
	assert.Equal(t, 8, len(strings_filled))
	for _, string_filled := range strings_filled{
		fmt.Println(string_filled)
		assert.NotContains(t, string_filled, "?")
	}
}
