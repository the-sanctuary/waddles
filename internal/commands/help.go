package commands

import (
	"fmt"
	"strings"

	"github.com/the-sanctuary/waddles/pkg/cmd"
)

var Help *cmd.Command = &cmd.Command{
	Name:        "help",
	Aliases:     []string{"h"},
	Description: "help text",
	Usage:       "help",
	Handler: func(c *cmd.Context) {
		c.ReplyStringf("Please use `%scommands` to view a list of commands.", c.Router.Config.Wadl.Prefix)
	},
}

var Commands *cmd.Command = &cmd.Command{
	Name:        "commands",
	Aliases:     []string{"c"},
	Description: "list of commands",
	Usage:       "commands",
	SubCommands: []*cmd.Command{},
	Handler: func(c *cmd.Context) {
		cmds := c.Router.Commands
		builder := strings.Builder{}

		builder.WriteString("```\n")

		rBuildHelp(c, &builder, cmds, 0)

		builder.WriteString("```")
		c.ReplyString(builder.String())
	},
}

func rBuildHelp(c *cmd.Context, builder *strings.Builder, cmds []*cmd.Command, depth int) {
	for _, cmd := range cmds {
		if cmd.HideInHelp {
			continue
		}
		indent := strings.Repeat("  ", depth)
		helpText := cmd.SPrintHelp()

		fmt.Fprintf(builder, "%s ♦ %s\n", indent, helpText)
		if cmd.HasSubcommands() {
			rBuildHelp(c, builder, cmd.SubCommands, depth+1)
		}
	}
}
