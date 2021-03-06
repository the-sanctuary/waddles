package commands

import "github.com/the-sanctuary/waddles/pkg/cmd"

var topicsManage *cmd.Command = &cmd.Command{
	Name:        "manage",
	Aliases:     []string{"m"},
	Description: "Manage all aspects of the topics system",
	Usage:       "manage (add|delete|edit|tags)",
	HideInHelp:  true,
	SubCommands: []*cmd.Command{topicsManageAdd, topicsManageDelete, topicsManageEdit, topicsManageTags},
	// Handler: func(c *cmd.Context) {}, //TODO: return basic stats about the topics system
}

var topicsManageAdd *cmd.Command = &cmd.Command{
	Name:        "add",
	Aliases:     []string{"a"},
	Description: "Add a topic",
	Usage:       "add <slug> <name> [<tag>[,<tag>,...]] <description>",
	Handler:     func(c *cmd.Context) {},
}

var topicsManageDelete *cmd.Command = &cmd.Command{
	Name:        "delete",
	Aliases:     []string{"d"},
	Description: "Delete the specified topic",
	Usage:       "delete <topic-slug>",
	Handler:     func(c *cmd.Context) {},
}

var topicsManageEdit *cmd.Command = &cmd.Command{
	Name:        "edit",
	Aliases:     []string{"e"},
	Description: "Edit a topic",
	Usage:       "edit <topic-slug> (slug|name|description|tags|archive)",
	Handler:     func(c *cmd.Context) {},
}
