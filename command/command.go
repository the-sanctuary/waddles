package command

import "github.com/bwmarrin/discordgo"

//Command  is the struct that holds information
type Command struct {
	Name        string
	Aliases     []string
	Description string
	Usage       string
	Example     string
	SubCommands []*Command
	Handler     Executor
}

type Context struct {
	//The current discordgo.Session
	Session *discordgo.Session
	//The message that started this execution
	Event *discordgo.MessageCreate
	//The command being executed
	Command *Command
}

//Executor represents an executor for a context execution
type Executor func(*Context)
