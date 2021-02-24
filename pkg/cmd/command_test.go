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

func Test_GeneratePermissionNode(t *testing.T) {
	testCmdSubSub := &Command{
		Name:    "subsub",
		Handler: func(c *Context) {},
	}

	testCmdSub := &Command{
		Name:        "sub",
		Handler:     func(c *Context) {},
		SubCommands: []*Command{testCmdSubSub},
	}

	testCmd := &Command{
		Name:        "testcmd",
		Handler:     func(c *Context) {},
		SubCommands: []*Command{testCmdSub},
	}

	nodes := testCmd.GeneratePermissionNode("")

	assert.Equal(t, "testcmd", nodes[0])
	assert.Equal(t, "testcmd.sub", nodes[1])
	assert.Equal(t, "testcmd.sub.subsub", nodes[2])
}
