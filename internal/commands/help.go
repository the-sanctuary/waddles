package commands

import (
	"strings"

	"github.com/the-sanctuary/waddles/pkg/cmd"
)

var help *cmd.Command = &cmd.Command{
	Name:        "help",
	Aliases:     []string{"h"},
	Description: "help text",
	Usage:       "help",
	Handler: func(c *cmd.Context) {
		c.ReplyStringf("Please use `%scommands` to view a list of commands.", c.Router.Config.Wadl.Prefix)
	},
}

var commands *cmd.Command = &cmd.Command{
	Name:        "commands",
	Aliases:     []string{"c"},
	Description: "list of commands",
	Usage:       "commands",
	SubCommands: []*cmd.Command{},
	Handler: func(c *cmd.Context) {
		cmds := c.Router.Commands
		builder := strings.Builder{}

		builder.WriteString("```\n")

		cmd.RBuildHelp(c, &builder, cmds, 0)

		builder.WriteString("```")
		c.ReplyString(builder.String())
	},
}
