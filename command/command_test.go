package command

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_BuildCommand(t *testing.T) {

	
	cmd := &Command{
		Name:        "test",
		Aliases:     []string{"test2"},
		Description: "this is a test command",
		Usage:       "test usage",
		Example:     "test example",
		SubCommands: []*Command{},
		Handler: func(c *Context) {
		},
	}

	assert.Equal(t, "test", cmd.Name)
	assert.Equal(t, []string{"test2"}, cmd.Aliases)
	assert.Equal(t, "this is a test command", cmd.Description)
	assert.Equal(t, "test usage", cmd.Usage)
	assert.Equal(t, "test example", cmd.Example)
	assert.Equal(t, "test", cmd.Handler)
	fmt.Println(cmd)
}
