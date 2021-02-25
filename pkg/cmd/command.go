package cmd

import (
	"fmt"

	"github.com/the-sanctuary/waddles/pkg/util"
)

//Command holds information about a command and what to do
type Command struct {
	//Name is the primary trigger and is used for each level of a permission node
	Name        string
	Aliases     []string
	Description string
	//Usage format: http://docopt.org/
	Usage string
	//HideInHelp whether or not this should be hidden from the help command.
	HideInHelp  bool
	SubCommands []*Command
	Handler     ContextExecutor
}

//Triggers returns all strings (the command name and any aliases) that trigger this command
func (c *Command) Triggers() []string {
	return append(c.Aliases, c.Name)
}

//HasSubcommands returns whether or not this command holds subcommands
func (c *Command) HasSubcommands() bool {
	if len(c.SubCommands) > 0 {
		return true
	}
	return false
}

//GeneratePermissionNode recursivly generates permission nodes for this and all subcommands and returns them
func (c *Command) GeneratePermissionNode(baseNode string) []string {
	nodes := make([]string, 0)

	newBaseNode := baseNode + c.Name
	nodes = append(nodes, newBaseNode)

	if c.HasSubcommands() {
		for _, subCmd := range c.SubCommands {
			newNodes := subCmd.GeneratePermissionNode(newBaseNode + ".")
			nodes = append(nodes, newNodes...)
		}
	}

	return nodes
}

//SPrintHelp returns the command's Usage and Description formatted in a nice way
func (c *Command) SPrintHelp() string {
	return fmt.Sprintf("%s - %s", c.Usage, c.Description)
}

//returns true and the command triggered by the provided string, otherwise returns (false, nil)
func triggerCheck(trigger string, cmds []*Command) (bool, *Command) {
	for _, cmd := range cmds {
		triggers := cmd.Triggers()

		if util.SliceContains(triggers, trigger) {
			return true, cmd
		}
	}
	return false, nil
}
