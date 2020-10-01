package command

import "strconv"

//Ping command
var ping *Command = &Command{
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
		if len(c.Args) >= 1 {
			n, err := strconv.Atoi(c.Args[0])

			if n > 5 {
				if c.Message.Author.ID == "90968241710563328" { //shame tim for being a shit
					c.Session.ChannelMessageSend(c.Message.ChannelID, "Bad boy Tim! That's too many pongs!!")
				} else {
					c.Session.ChannelMessageSend(c.Message.ChannelID, "That's too many!")
				}
				return
			}

			if err != nil {
				// TODO: Print an error, count must be a number
			}
			for i := 0; i < n; i++ {
				c.Session.ChannelMessageSend(c.Message.ChannelID, "Pong!")
			}
		} // else error

	},
}
