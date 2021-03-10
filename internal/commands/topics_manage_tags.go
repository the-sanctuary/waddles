package commands

import "github.com/the-sanctuary/waddles/pkg/cmd"

var topicsManageTags *cmd.Command = &cmd.Command{
	Name:        "tags",
	Aliases:     []string{"t"},
	Description: "Manage Tags",
	Usage:       "tags (list|add|edit|delete)",
	Handler:     func(c *cmd.Context) {},
	SubCommands: []*cmd.Command{topicsManageTagsEdit},
}

var topicsManageTagsEdit *cmd.Command = &cmd.Command{
	Name:        "edit",
	Aliases:     []string{"e"},
	Description: "Edit a tag",
	Usage:       "edit <tag-id> (name|description) <new value>",
	Handler:     func(c *cmd.Context) {},
}
