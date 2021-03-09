package commands

import (
	"strings"

	"github.com/the-sanctuary/waddles/pkg/cmd"
	"github.com/the-sanctuary/waddles/pkg/db"
	"gorm.io/gorm"
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
		tags := strings.Split(c.Args[2], ",")
		description := strings.Join(c.Args[3:], " ")

		topic := db.Topic{
			Slug:        slug,
			Name:        name,
			Description: description,
			Tags:        tagsFromSlice(c.DB(), tags),
		}

		tx := db.Instance.Create(&topic)

		if tx.Error == nil {
			c.ReplyStringf("Topic (%d) Created!", topic.ID)
		} else {
			c.ReplyError(tx.Error)
		}
	},
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
	Handler: func(c *cmd.Context) {
		if len(c.Args) < 3 {
			c.ReplyString("Invalid synax.")
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

		target := c.Args[1]

		switch target {
		case "slug":
			newSlug := c.Args[2]

			_, err := db.TopicFindBySlug(c.DB(), newSlug)
			if err == gorm.ErrRecordNotFound {
				topic.Slug = newSlug
			} else if err != nil {
				c.ReplyError(err)
			} else {
				c.ReplyStringf("A topic already is using the slug: `%s`", newSlug)
			}
		case "name":
			topic.Name = c.Args[2]
		case "description":
			topic.Description = strings.Join(c.Args[2:], " ")
		case "tags":
			//TODO: this should add/remove; not override completely.
			// It instead should look like this: tags (add|delete) <tag>[,<tag>,...]
			topic.Tags = tagsFromSlice(c.DB(), strings.Split(c.Args[3], ","))
		case "archive":
			topic.Archived = true
		default:
			c.ReplyStringf("Invalid target: `%v`", target)
			return
		}

		tx := c.DB().Save(&topic)
		if tx.Error != nil {
			c.ReplyError(tx.Error)
			return
		}

		c.ReplyStringf("Successfully updated value: `%s`", target)
	},
}

func tagsFromSlice(wdb *db.WadlDB, tags []string) []*db.TopicTag {
	var topicTags []*db.TopicTag

	for _, tag := range tags {
		topicTag := &db.TopicTag{Name: tag}

		wdb.Where(topicTag).FirstOrCreate(&topicTag)

		topicTags = append(topicTags, topicTag)
	}

	return topicTags
}
