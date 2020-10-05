package command

var nitro *Command = &Command{
	Name:        "nitro",
	Aliases:     *&[]string{"n"},
	Description: " access your perks as a server nitro booster",
	Usage:       "nitro ",
	SubCommands: []*Command{nitroChannel},
	Handler: func(c *Context) {
		c.ReplyString("nitro")
	},
}
