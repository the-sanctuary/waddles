package commands

import "github.com/the-sanctuary/waddles/pkg/cmd"

//Commands returns all the waddles commands
func Commands() []*cmd.Command {
	return []*cmd.Command{
		ping,
		purge,
		uptime,
		nitro,
		debug,
		helpCommands,
		help,
	}
}
