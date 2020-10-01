package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FindDeepestHandler(t *testing.T) {
	testcmdSub := &Command{
		Name:    "sub",
		Handler: func(c *Context) {},
	}

	testcmd := &Command{
		Name:        "testcmd",
		Handler:     func(c *Context) {},
		SubCommands: []*Command{testcmdSub},
	}

	actual, args := findDeepestCommand(testcmd, []string{"testcmd", "sub", "Arrrrrrghhhh"})
	assert.Equal(t, actual.Name, "sub")
	assert.Equal(t, args, []string{"Arrrrrrghhhh"})
}
