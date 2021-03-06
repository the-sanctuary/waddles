package commands

import "github.com/the-sanctuary/waddles/pkg/cmd"

var topics *cmd.Command = &cmd.Command{
	Name:        "topics",
	Aliases:     []string{""},
	Description: "View/Manage your topic subscriptions on this server",
	Usage:       "topics (add|list|remove)",
	SubCommands: []*cmd.Command{topicsList, topicsAdd, topicsRemove, topicsManage},
}

var topicsList *cmd.Command = &cmd.Command{
	Name:        "list",
	Aliases:     []string{"l"},
	Description: "List all the available topics",
	Usage:       "list",
	Handler:     func(c *cmd.Context) {},
}
var topicsAdd *cmd.Command = &cmd.Command{
	Name:        "add",
	Aliases:     []string{"a"},
	Description: "Subscribe yourself to a topic",
	Usage:       "add <topic-slug>",
	Handler:     func(c *cmd.Context) {},
}
var topicsRemove *cmd.Command = &cmd.Command{
	Name:        "remove",
	Aliases:     []string{"a"},
	Description: "Unsubscribe from a topic.",
	Usage:       "remove <topic-slug>",
	Handler:     func(c *cmd.Context) {},
}
