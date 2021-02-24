package commands

import (
	"github.com/the-sanctuary/waddles/pkg/cmd"
)

var nitro *cmd.Command = &cmd.Command{
	Name:        "nitro",
	Aliases:     []string{"n"},
	Description: " access your perks as a server nitro booster",
	Usage:       "nitro (channel)",
	SubCommands: []*cmd.Command{nitroChannel},
	Handler: func(c *cmd.Context) {
		c.ReplyHelp()
	},
}
