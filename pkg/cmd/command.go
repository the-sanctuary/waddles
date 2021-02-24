package cmd

import (
	"fmt"
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

//SPrintHelp returns the formatted help text string
func (c *Command) SPrintHelp() string {
	return fmt.Sprintf("%s - %s", c.Usage, c.Description)
}
