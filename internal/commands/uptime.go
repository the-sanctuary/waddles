package commands

import (
	"time"

	"github.com/the-sanctuary/waddles/pkg/cmd"

	"github.com/the-sanctuary/waddles/pkg/util"
)

var uptime *cmd.Command = &cmd.Command{
	Name:        "uptime",
	Description: "the uptime of the bot",
	Usage:       "uptime",
	Handler: func(c *cmd.Context) {
		c.ReplyStringf("Current Bot Time: `%s`\nUptime: `%s`", time.Now().Format(time.RFC1123Z), util.Uptime().Round(time.Second))
	},
}
