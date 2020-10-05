package command

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/the-sanctuary/waddles/util"
	"gorm.io/gorm"
)

//Command  is the struct that holds information about a command
type Command struct {
	Name        string
	Aliases     []string
	Description string
	//Usage format: http://docopt.org/
	Usage       string
	SubCommands []*Command
	Handler     ContextExecutor
}

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

//Triggers returns all strings (the command name and any aliases) that trigger this command
func (c *Command) Triggers() []string {
	return append(c.Aliases, c.Name)
}

//DB returns the current gorm.DB instance
func (ctx *Context) DB() *gorm.DB {
	return ctx.Router.WadlDB.DB
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
	return ctx.ReplyStringf("%s %s %s", ctx.Router.Prefix, ctx.Command.Name, ctx.Command.Usage)
}

func (ctx *Context) ReplyError(err error) bool {
	if err != nil {
		ctx.ReplyString("An error occured. Check the log for details.")
		util.DebugError(err)
		return true
	}
	return false
}

func (c *Command) hasSubcommands() bool {
	if len(c.SubCommands) > 0 {
		return true
	}
	return false
}
