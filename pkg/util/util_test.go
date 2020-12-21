package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SliceContains(t *testing.T) {
	slice := []string{"string1", "string@2!", "pumpkins", "6420", "!@#$%^"}
	assert.True(t, SliceContains(slice, "string@2!"))
	assert.True(t, SliceContains(slice, "pumpkins"))
	assert.False(t, SliceContains(slice, "pump"))
	assert.False(t, SliceContains(slice, "string"))
}
