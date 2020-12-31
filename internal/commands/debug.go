package commands

import (
	"fmt"
	"strings"

	"github.com/the-sanctuary/waddles/pkg/cmd"
)

//Debug debug command
var Debug *cmd.Command = &cmd.Command{
	Name:        "debug",
	Aliases:     []string{"dbg"},
	Description: "bot debug interface",
	Usage:       "",
	HideInHelp:  true,
	SubCommands: []*cmd.Command{debugDumpPerms, debugListPerms},
	Handler: func(c *cmd.Context) {
		c.ReplyHelp()
	},
}

var debugListPerms *cmd.Command = &cmd.Command{
	Name:        "listPerms",
	Aliases:     []string{""},
	Description: "lists text representation of permission system",
	Usage:       "",
	Handler: func(c *cmd.Context) {
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

var debugDumpPerms *cmd.Command = &cmd.Command{
	Name:        "dumpPerms",
	Aliases:     []string{""},
	Description: "dumps raw PermissionSystem struct",
	Usage:       "",
	Handler: func(c *cmd.Context) {
		tre := &c.Router.PermSystem.Tree
		c.ReplyStringf("```%+v```", tre)
	},
}
