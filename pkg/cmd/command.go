package cmd

import (
	"fmt"

	"github.com/the-sanctuary/waddles/pkg/permissions"
)

//Command  is the struct that holds information about a command
type Command struct {
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

//GeneratePermissionNode recursivly adds a  permission node to this permission system
func (c *Command) GeneratePermissionNode(permSystem *permissions.PermissionSystem, baseNode string) {
	newBaseNode := baseNode + c.Name

	permSystem.AddPermissionNode(newBaseNode)

	if c.HasSubcommands() {
		for _, subCmd := range c.SubCommands {
			subCmd.GeneratePermissionNode(permSystem, newBaseNode+".")
		}
	}
}

//SPrintHelp returns the formatted help text string
func (c *Command) SPrintHelp() string {
	return fmt.Sprintf("%s - %s", c.Usage, c.Description)
}
