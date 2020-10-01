package command

//Ping command
var Ping *Command = &Command{
	Name:        "ping",
	Aliases:     *&[]string{"pong"},
	Description: "This pongs your ping(pong)!",
	Usage:       "ping",
	Handler: func(c *Context) {
		c.Session.ChannelMessageSend(c.Message.ChannelID, "Pong!")
	},
	SubCommands: []*Command{pingCount},
}

var pingCount *Command = &Command{
	Name:        "count",
	Description: "how many times to reply with pong",
	Usage:       "ping",
	Handler: func(c *Context) {
		// c.Session.ChannelMessageSend(c.Event.Channel.ID, "Pong!")
	},
}
