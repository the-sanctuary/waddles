package command

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/db"
	"github.com/the-sanctuary/waddles/util"
)

//Router is the central command multiplexer
type Router struct {
	Commands []*Command
	Prefix   string
	WadlDB   *db.WadlDB
}

//BuildRouter returns a fully built router stuct with commands preregistered
func BuildRouter(wdb *db.WadlDB) Router {
	r := Router{
		Prefix: util.Cfg.Wadl.Prefix,
		WadlDB: wdb,
	}
	r.RegisterCommands(
		ping,
		purge,
		uptime,
		nitro,
	)
	return r
}

//RegisterCommands adds a command(s) to the Router
func (r *Router) RegisterCommands(cmds ...*Command) {
	for _, c := range cmds {
		r.Commands = append(r.Commands, c)
	}
}

//Handler returns the func that deals with command delegates execution to command
func (r *Router) Handler() func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(session *discordgo.Session, message *discordgo.MessageCreate) {
		log.Trace().Msg("Entering Router Handler")
		defer log.Trace().Msg("Exiting Router Handler")

		// Check if the message was sent by a bot
		if message.Author.Bot {
			log.Debug().Msgf("Ignoring message from bot: %s (%s)", message.Author.Username, message.Author.ID)
			return
		}

		// Check to see if we're being pinged
		if message.Content == fmt.Sprintf("<@!%s>", session.State.User.ID) {
			log.Debug().Msgf("Received ping from: %s", session.State.User.Username)
			handlePing(session, message)
			return
		}

		// Check to see if our prefix is there
		if message.Content[:len(r.Prefix)] != r.Prefix {
			return
		}

		split := strings.Split(message.Content[len(r.Prefix):], " ")
		correct, cmd := triggerCheck(split[0], r.Commands)

		if correct {
			cmd, args := findDeepestCommand(cmd, split)
			ctx := buildContext(session, message, cmd, args, r)
			cmd.Handler(&ctx)

			//Update UserActivity entry's CommandCount
			var ua db.UserActivity
			r := db.CurrentWadlDB().DB.Where(&db.UserActivity{UserID: message.Author.ID}).FirstOrCreate(&ua)
			util.DebugError(r.Error)
			ua.CommandCount++
			db.CurrentWadlDB().DB.Save(&ua)
		}
	}
}

//Finds and returns the deepest subcommand for a given command and arg slice
func findDeepestCommand(prevCmd *Command, args []string) (*Command, []string) {
	if len(prevCmd.SubCommands) > 0 {
		if len(args) > 1 {
			found, cmd := triggerCheck(args[1], prevCmd.SubCommands)
			if found {
				return findDeepestCommand(cmd, args[1:])
			}
		}
	}
	return prevCmd, args[1:]
}

//returns the command triggered by the provided string, otherwise returns (false, nil)
func triggerCheck(trigger string, cmds []*Command) (bool, *Command) {
	for _, cmd := range cmds {
		triggers := cmd.Triggers()

		if util.SliceContains(triggers, trigger) {
			return true, cmd
		}
	}
	return false, nil
}

func buildContext(session *discordgo.Session, message *discordgo.MessageCreate, command *Command, args []string, router *Router) Context {
	return *&Context{
		Session: session,
		Message: message,
		Command: command,
		Args:    args,
		Router:  router,
	}
}

func handlePing(session *discordgo.Session, message *discordgo.MessageCreate) {}
