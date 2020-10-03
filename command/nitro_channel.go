package command

var nitroChannel *Command = &Command{
	Name:        "nbperks",
	Aliases:     *&[]string{"nbp"},
	Description: " access your perks as a server nitro booster",
	Usage:       "nbp []",
	Handler: func(c *Context) {
	},
}
