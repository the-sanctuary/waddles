package commands

import (
	"fmt"
	"strings"

	"github.com/the-sanctuary/waddles/pkg/cmd"
	"github.com/the-sanctuary/waddles/pkg/db"
)

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
	Handler: func(c *cmd.Context) {
		allTopics := db.TopicFindAll(c.DB())

		if len(allTopics) < 1 {
			c.ReplyString("There are no topics.")
			return
		}

		builder := strings.Builder{}

		for _, topic := range allTopics {
			builder.WriteString(fmt.Sprintf("%s (%s) - %s \n", topic.Name, topic.Slug, topic.Description))
		}

		c.ReplyString(builder.String())
	},
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
