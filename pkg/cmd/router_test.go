package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/the-sanctuary/waddles/pkg/permissions"
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

	actual, args, node := findDeepestCommand(testcmd, []string{"testcmd", "sub", "Arrrrrrghhhh"}, testcmd.Name)
	assert.Equal(t, "sub", actual.Name)
	assert.Equal(t, []string{"Arrrrrrghhhh"}, args)
	assert.Equal(t, "testcmd.sub", node)
}

func Test_GeneratePermissionNodes(t *testing.T) {
	testCmdSub1 := &Command{
		Name:    "sub1",
		Handler: func(c *Context) {},
	}

	testCmdSub21 := &Command{
		Name:    "sub2-1",
		Handler: func(c *Context) {},
	}

	testCmdSub2 := &Command{
		Name:        "sub2",
		Handler:     func(c *Context) {},
		SubCommands: []*Command{testCmdSub21},
	}

	testCmd := &Command{
		Name:        "testcmd",
		Handler:     func(c *Context) {},
		SubCommands: []*Command{testCmdSub1, testCmdSub2},
	}

	r := Router{
		PermSystem: &permissions.PermissionSystem{
			Nodes: make([]*permissions.PermissionNode, 0),
		},
	}

	r.RegisterCommands([]*Command{testCmd})
	r.generatePermissionNodes()

	generatedNodes := r.PermSystem.Nodes

	assert.Equal(t, "testcmd", generatedNodes[0].Identifier)
	assert.Equal(t, "testcmd.sub1", generatedNodes[1].Identifier)
	assert.Equal(t, "testcmd.sub2", generatedNodes[2].Identifier)
	assert.Equal(t, "testcmd.sub2.sub2-1", generatedNodes[3].Identifier)
}
