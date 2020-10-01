package command

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/util"
)

//Router is the central command multiplexer
type router struct {
	Commands []*Command
	Prefix   string
}

func BuildRouter() router {
	r := router{
		Prefix: "~",
	}
	r.RegisterCommands(
		ping,
	)
	return r
}

//RegisterCommand adds a command to the Router
func (r *router) RegisterCommands(cmds ...*Command) {
	for _, c := range cmds {
		r.Commands = append(r.Commands, c)
	}
}

//Handler returns the func that deals with command delegates execution to command
func (r *router) Handler() func(*discordgo.Session, *discordgo.MessageCreate) {
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

		//working setup for base commands
		// trigger := strings.SplitN(message.Content, " ", 2)[0][1:]
		// correct, cmd := triggerCheck(trigger, r.Commands)
		// args := strings.Split(message.Content, " ")

		// log.Trace().Msgf("trigger: %s, correct: %t, cmd: %s", trigger, correct, cmd.Name)

		split := strings.Split(message.Content[len(r.Prefix):], " ")
		correct, cmd := triggerCheck(split[0], r.Commands)

		if correct {
			cmd, args := findDeepestCommand(cmd, split)
			ctx := buildContext(session, message, cmd, args)
			cmd.Handler(&ctx)
		}
	}
}

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

func triggerCheck(trigger string, cmds []*Command) (bool, *Command) {
	for _, cmd := range cmds {
		triggers := cmd.Triggers()

		// log.Trace().Msgf("Tested for cmd trigger: %s in: [%s]", testCmd, triggers)

		if util.SliceContains(triggers, trigger) {
			return true, cmd
		}
	}
	return false, nil
}

func buildContext(session *discordgo.Session, message *discordgo.MessageCreate, command *Command, args []string) Context {
	return *&Context{
		Session: session,
		Message: message,
		Command: command,
		Args:    args,
	}
}

func handlePing(session *discordgo.Session, message *discordgo.MessageCreate) {}
