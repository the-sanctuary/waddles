package command

import (
	"strconv"
	"time"

	"github.com/the-sanctuary/waddles/util"
)

//Purge command
var purge *Command = &Command{
	Name:        "purge",
	Description: "Remove message history.",
	Usage:       "purge [num]",
	Handler: func(c *Context) {
		if !(util.SliceContains(c.Message.Member.Roles, "244943928913166338") || util.SliceContains(c.Message.Member.Roles, "183808574227611649")) { // TODO: permissions system
			c.ReplyString("You must be a Lord of the Server or Royal Moderator to use this command.")
			return
		}

		if len(c.Args) >= 1 {
			n, err := strconv.Atoi(c.Args[0])
			if util.DebugError(err) {
				c.ReplyString("An error occured. Check the log for details.")
				return
			}

			//make sure n is positve or not zero
			if n < 1 {
				c.ReplyString("You must delete at least one message.")
			}

			messages, err := c.Session.ChannelMessages(c.Message.ChannelID, n+1, "", "", "")
			if util.DebugError(err) {
				c.ReplyString("An error occured. Check the log for details.")
			}

			var msgIds []string
			for _, m := range messages {
				msgIds = append(msgIds, m.ID)
			}

			if len(msgIds) == 1 {
				c.Session.ChannelMessageDelete(c.Message.ChannelID, msgIds[0])
			} else {
				c.Session.ChannelMessagesBulkDelete(c.Message.ChannelID, msgIds)
			}

			msg := c.ReplyStringf("Deleted %d messages from this channel.", n)

			go func() {
				time.Sleep(3000 * time.Millisecond)
				c.Session.ChannelMessageDelete(c.Message.ChannelID, msg.ID)
			}()
		} else {
			c.ReplyString("You must supply the number of messages to purge from the current channel")
		}
	},
}
