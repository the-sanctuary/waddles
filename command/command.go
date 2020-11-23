package command

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
