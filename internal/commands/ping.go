package commands

import (
	"strconv"

	"github.com/the-sanctuary/waddles/pkg/cmd"

	"github.com/the-sanctuary/waddles/pkg/util"
)

//Ping command
var Ping *cmd.Command = &cmd.Command{
	Name:        "ping",
	Aliases:     *&[]string{"pong"},
	Description: "This pongs your ping(pong)!",
	Usage:       "ping [count <num>]",
	Handler: func(c *cmd.Context) {
		c.ReplyString("Pong!")
	},
	SubCommands: []*cmd.Command{PingCount},
}

var PingCount *cmd.Command = &cmd.Command{
	Name:        "count",
	Description: "how many times to reply with pong",
	Usage:       "count <num>",
	Handler: func(c *cmd.Context) {
		if len(c.Args) >= 1 {
			n, err := strconv.Atoi(c.Args[0])

			if util.DebugError(err) {
				c.ReplyString("The argument to count must be a postive integer")
				return
			}

			//make sure n is a postive number
			n = util.AbsInt(n)

			if n > 5 {
				if c.Message.Author.ID == "90968241710563328" { //shame tim for being a shit
					c.ReplyString("Bad boy Tim! That's too many pongs!!")
				} else {
					c.ReplyString("That's too many!")
				}
				return
			}

			for i := 0; i < n; i++ {
				c.ReplyString("Pong!")
			}
		} else {
			c.ReplyString("`count` subcommand must have an arguement supplied.")
			c.ReplyHelp()
		}
	},
}
