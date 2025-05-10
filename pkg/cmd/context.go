package cmd

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/the-sanctuary/waddles/pkg/db"
	"github.com/the-sanctuary/waddles/pkg/util"
)

//ContextExecutor represents an executor for a context execution
type ContextExecutor func(*Context)

//Context holds all state for a command's execution
type Context struct {
	//The current discordgo.Session
	Session *discordgo.Session
	//The message that started this execution
	Message *discordgo.MessageCreate
	//The command being executed
	Command *Command
	//Command args (i.e. the split message content)
	Args []string
	//The router currently used
	Router *Router
}

//DB returns the current db.WadlDB instance
func (ctx *Context) DB() *db.WadlDB {
	return ctx.Router.WadlDB
}

//ReplyString replies to the contextual channel with the string provided.
// returns nil if an error occured while sending the message
func (ctx *Context) ReplyString(message string) *discordgo.Message {
	msg, err := ctx.Session.ChannelMessageSend(ctx.Message.ChannelID, message)
	if util.DebugError(err) {
		return nil
	}
	return msg
}

//ReplyStringf replies to the contextual channel with the string provided
func (ctx *Context) ReplyStringf(format string, a ...interface{}) *discordgo.Message {
	return ctx.ReplyString(fmt.Sprintf(format, a...))
}

//ReplyHelp prints the command's help text to the provided Context
func (ctx *Context) ReplyHelp() *discordgo.Message {
	return ctx.ReplyStringf("Usage: `%s`", ctx.Command.SPrintHelp())
}

//ReplyError returns a generic error to the user
func (ctx *Context) ReplyError(err error) bool {
	if err != nil {
		//TODO: if user is a debugUser, send them a PM with the error itself instead of just a generic response to the channel.
		// authorID := ctx.Message.Author.ID
		// if util.SliceContains(cfg.ReadConfig().Permissions.DebugUsers, authorID) {
		// 	st, err := ctx.Session.UserChannelCreate(authorID)
		// 	util.DebugError(err)
		// 	ctx.Session.ChannelMessageSend(st.ID, fmt.Sprintf("Error Report: ```%s```", err.Error()))
		// }
		ctx.ReplyString("An error occured. Check the log for details.")
		util.DebugError(err)
		return true
	}
	return false
}

func (ctx *Context) ReplyTimeDeleteStringf(delay time.Duration, format string, a ...interface{}) {
	errMsg := ctx.ReplyStringf(format, a...)

	time.Sleep(delay)

	ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, errMsg.ID)
}
