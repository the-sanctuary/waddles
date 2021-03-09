package commands

import (
	"fmt"
	"strings"

	"github.com/the-sanctuary/waddles/pkg/cmd"
	"github.com/the-sanctuary/waddles/pkg/db"
	"gorm.io/gorm"
)

var topics *cmd.Command = &cmd.Command{
	Name:        "topics",
	Aliases:     []string{""},
	Description: "View/Manage your topic subscriptions on this server",
	Usage:       "topics (add|list|remove)",
	SubCommands: []*cmd.Command{topicsList, topicsAdd, topicsSubscribed, topicsRemove, topicsManage},
}

var topicsList *cmd.Command = &cmd.Command{
	Name:        "list",
	Aliases:     []string{"l"},
	Description: "List all the available topics",
	Usage:       "list",
	Handler: func(c *cmd.Context) {
		allTopics, err := db.TopicFindAll(c.DB())
		if err != nil {
			c.ReplyError(err)
			return
		}

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

var topicsSubscribed *cmd.Command = &cmd.Command{
	Name:        "subscribed",
	Aliases:     []string{"s"},
	Description: "List your subscribed topics",
	Usage:       "subscribed",
	Handler: func(c *cmd.Context) {
		allTopics, err := db.TopicFindAllForUser(c.DB(), &db.User{DiscordID: c.Message.Author.ID})
		if err != nil {
			c.ReplyError(err)
			return
		}

		if len(allTopics) < 1 {
			c.ReplyString("You aren't subscribed to any topics.")
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
	Handler: func(c *cmd.Context) {
		if len(c.Args) != 1 {
			c.ReplyString("Error: You should only have argument.")
			return
		}
		slug := c.Args[0]

		topic, err := db.TopicFindBySlug(c.DB(), slug)
		if err == gorm.ErrRecordNotFound {
			c.ReplyStringf("Error: Couldn't find topic: `%s`", slug)
			return
		} else if err != nil {
			c.ReplyError(err)
			return
		}

		var topicUser db.TopicUser
		tx := c.DB().Where("topic_id = ?", topic.ID).Where("discord_id = ?", c.Message.Author.ID).First(&topicUser)

		if tx.Error == gorm.ErrRecordNotFound {
			tx2 := c.DB().Create(&db.TopicUser{User: db.User{DiscordID: c.Message.Author.ID}, Topic: topic, Active: true})
			if tx2.Error != nil {
				c.ReplyError(tx2.Error)
				return
			}
		} else if tx.Error != nil {
			c.ReplyError(tx.Error)
			return
		}

		if topicUser.Active {
			c.ReplyStringf("You are already subscribed to `%s (%s)`.", topic.Name, topic.Slug)
			return
		}
		topicUser.Active = true

		c.DB().Save(&topicUser)
		c.ReplyStringf("You have been subscribed to `%s (%s)`.", topic.Name, topic.Slug)
	},
}
var topicsRemove *cmd.Command = &cmd.Command{
	Name:        "remove",
	Aliases:     []string{"r"},
	Description: "Unsubscribe from a topic.",
	Usage:       "remove <topic-slug>",
	Handler: func(c *cmd.Context) {
		if len(c.Args) != 1 {
			c.ReplyString("Error: You should only have argument.")
			return
		}

		slug := c.Args[0]

		topic, err := db.TopicFindBySlug(c.DB(), slug)
		if err == gorm.ErrRecordNotFound {
			c.ReplyStringf("Error: Couldn't find topic: `%s`", slug)
			return
		} else if err != nil {
			c.ReplyError(err)
			return
		}

		var topicUser db.TopicUser
		tx := c.DB().Where("topic_id = ?", topic.ID).Where("discord_id = ?", c.Message.Author.ID).Where("active = ?", true).First(&topicUser)

		if tx.Error == gorm.ErrRecordNotFound {
			c.ReplyStringf("You weren't subscribed that topic anyway.")
			return
		} else if err != nil {
			c.ReplyError(err)
			return
		}

		topicUser.Active = false

		c.DB().Save(&topicUser)
		c.ReplyStringf("You have been unsubscribed from topic: `%s (%s)`", topic.Name, topic.Slug)
	},
}
