package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Triggers(t *testing.T) {
	cmd := &Command{
		Name:        "test",
		Aliases:     []string{"test2"},
		Description: "this is a test command",
		Usage:       "test usage",
		SubCommands: []*Command{},
		Handler:     func(c *Context) {},
	}

	assert.Equal(t, []string{"test2", "test"}, cmd.Triggers())
}
