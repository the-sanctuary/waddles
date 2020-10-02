package command

import "strconv"

//Purge command
var purge *Command = &Command{
	Name:        "purge",
	Description: "Remove message history.",
	Usage:       "purge 10",
	Handler: func(c *Context) {
		if len(c.Args) >= 1 {
			n, err := strconv.Atoi(c.Args[0])
			if err != nil {
				// TODO: Print an error, count must be a number
			}
			if n < 0 {
				n = 0
			}
			messages, err := c.Session.ChannelMessages(c.Message.ChannelID, n+1, "", "", "")
			if err != nil {
				// TODO: Print an error, unable to get messages
			}
			var msgIds []string
			for _, m := range messages {
				msgIds = append(msgIds, m.ID)
			}
			if len(msgIds) == 0 {
				// TODO: Must delete at least 1 message
			} else if len(msgIds) == 1 {
				c.Session.ChannelMessageDelete(c.Message.ChannelID, msgIds[0])
			} else {
				c.Session.ChannelMessagesBulkDelete(c.Message.ChannelID, msgIds)
			}
		} // else error
	},
}
