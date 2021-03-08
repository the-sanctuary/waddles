package commands

import (
	"strings"

	"github.com/the-sanctuary/waddles/pkg/cmd"
	"github.com/the-sanctuary/waddles/pkg/db"
)

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
	Handler: func(c *cmd.Context) {
		if len(c.Args) < 4 {
			c.ReplyString("You haven't entered enough arguments.")
			return
		}

		slug := c.Args[0]
		name := c.Args[1]
		tags := c.Args[2]
		description := c.Args[3:]

		topic := db.Topic{
			Slug:        slug,
			Name:        name,
			Description: strings.Join(description, " "),
			Tags:        tagsFromSlice(strings.Split(tags, ",")),
		}

		tx := db.Instance.Create(&topic)

		if tx.Error == nil {
			c.ReplyStringf("Topic (%d) Created!", topic.ID)
		} else {
			c.ReplyError(tx.Error)
		}
	},
}

func tagsFromSlice(tags []string) []*db.TopicTag {
	var topicTags []*db.TopicTag

	for _, tag := range tags {
		topicTags = append(topicTags, &db.TopicTag{Name: tag})
	}

	return topicTags
}

var topicsManageDelete *cmd.Command = &cmd.Command{
	Name:        "delete",
	Aliases:     []string{"d"},
	Description: "Delete the specified topic",
	Usage:       "delete <topic-slug>",
	Handler: func(c *cmd.Context) {
		if len(c.Args) != 1 {
			c.ReplyString("You must only supply the topic slug.")
			return
		}

		slug := c.Args[0]

		var topic db.Topic

		tx := c.DB().Where("slug = ?", slug).First(&topic)
		if tx.Error != nil {
			c.ReplyStringf("Topic with given slug: `%s` not found.", slug)
			c.ReplyError(tx.Error)
			return
		}

		tx = c.DB().Delete(&topic)

		if tx.Error == nil {
			c.ReplyStringf("Topic %s (%s) deleted!", topic.Name, topic.Slug)
		} else {
			c.ReplyError(tx.Error)
		}
	},
}

var topicsManageEdit *cmd.Command = &cmd.Command{
	Name:        "edit",
	Aliases:     []string{"e"},
	Description: "Edit a topic",
	Usage:       "edit <topic-slug> (slug|name|description|tags|archive)",
	Handler:     func(c *cmd.Context) {},
}
