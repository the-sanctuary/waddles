package command

import (
	"time"

	"github.com/the-sanctuary/waddles/util"
)

var uptime *Command = &Command{
	Name:        "uptime",
	Description: "the uptime of the bot",
	Usage:       "uptime",
	Handler: func(c *Context) {
		c.ReplyStringf("Current Bot Time: `%s`", time.Now().Format(time.RFC1123Z))
		c.ReplyStringf("Uptime: `%s`", util.Uptime().Round(time.Second))
	},
}
