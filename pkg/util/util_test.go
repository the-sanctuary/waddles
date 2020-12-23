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

func Test_AbsInt(t *testing.T) {
	assert.Equal(t, 123, AbsInt(123))
	assert.Equal(t, 123, AbsInt(-123))
	assert.Equal(t, 0, AbsInt(0))
	assert.Equal(t, 0, AbsInt(-0))
}

func Test_FileExists(t *testing.T) {
	assert.FileExists(t, "./util.go")
	assert.True(t, FileExists("./util.go"))

	assert.NoFileExists(t, "./fake_util.go")
	assert.False(t, FileExists("./fake_util.go"))

	assert.False(t, FileExists("not a path at all"))
}
