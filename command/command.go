package command

import "github.com/bwmarrin/discordgo"

//Command  is the struct that holds information about a command
type Command struct {
	Name        string
	Aliases     []string
	Description string
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
}

//Triggers returns all strings (the command name and any aliases) that trigger this command
func (c *Command) Triggers() []string {
	return append(c.Aliases, c.Name)
}

func (c *Command) hasSubcommands() bool {
	if len(c.SubCommands) > 0 {
		return true
	}
	return false
}
