package commands

import "github.com/the-sanctuary/waddles/pkg/command"

var Nitro *command.Command = &command.Command{
	Name:        "nitro",
	Aliases:     []string{"n"},
	Description: " access your perks as a server nitro booster",
	Usage:       "nitro channel ",
	SubCommands: []*command.Command{NitroChannel},
	Handler: func(c *command.Context) {
		c.ReplyHelp()
	},
}
