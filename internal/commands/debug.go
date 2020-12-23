package commands

import (
	"fmt"
	"strings"

	"github.com/the-sanctuary/waddles/pkg/command"
)

var Debug *command.Command = &command.Command{
	Name:        "debug",
	Aliases:     []string{"dbg"},
	Description: "bot debug interface",
	Usage:       "",
	SubCommands: []*command.Command{DebugDumpPerms, DebugListPerms},
	Handler: func(c *command.Context) {
		c.ReplyHelp()
	},
}

var DebugListPerms *command.Command = &command.Command{
	Name:        "listPerms",
	Aliases:     []string{""},
	Description: "lists text representation of permission system",
	Usage:       "",
	Handler: func(c *command.Context) {
		tre := &c.Router.PermSystem.Tree

		var lines []string

		lines = append(lines, "`===Nodes===`")
		for _, node := range c.Router.PermSystem.Nodes {
			lines = append(lines, " - "+node.Identifier)
		}

		lines = append(lines, " ")

		lines = append(lines, "`===Sets===`")
		for _, set := range tre.Sets {
			lines = append(lines, "Name: "+set.Name)
			lines = append(lines, " - Description: "+set.Description)
			lines = append(lines, " - Nodes: ")
			for _, node := range set.Nodes {
				lines = append(lines, "   - "+node.Identifier)
			}
		}

		lines = append(lines, " ")

		lines = append(lines, "`===Groups===`")
		for _, group := range tre.Groups {
			lines = append(lines, "Name: "+group.Name)
			lines = append(lines, "  - Description: "+group.Description)
			lines = append(lines, fmt.Sprintf("  - Role: <@&%s> (%s)", group.RoleID, group.RoleID))

			lines = append(lines, "  - Sets: ")
			for _, set := range group.Sets {
				lines = append(lines, "   - "+set.Name)
			}
		}

		lines = append(lines)

		c.ReplyString(strings.Join(lines, "\n"))
	},
}

var DebugDumpPerms *command.Command = &command.Command{
	Name:        "dumpPerms",
	Aliases:     []string{""},
	Description: "dumps raw PermissionSystem struct",
	Usage:       "",
	Handler: func(c *command.Context) {
		tre := &c.Router.PermSystem.Tree
		c.ReplyStringf("```%+v```", tre)
	},
}
