package commands

import (
	"time"

	"github.com/the-sanctuary/waddles/pkg/command"
	"github.com/the-sanctuary/waddles/pkg/util"
)

var Uptime *command.Command = &command.Command{
	Name:        "uptime",
	Description: "the uptime of the bot",
	Usage:       "uptime",
	Handler: func(c *command.Context) {
		c.ReplyStringf("Current Bot Time: `%s`", time.Now().Format(time.RFC1123Z))
		c.ReplyStringf("Uptime: `%s`", util.Uptime().Round(time.Second))
	},
}
