package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/internal/model"
	"github.com/the-sanctuary/waddles/pkg/cfg"
	"github.com/the-sanctuary/waddles/pkg/db"
	"github.com/the-sanctuary/waddles/pkg/permissions"
	"github.com/the-sanctuary/waddles/pkg/util"
)

//Router is the central command multiplexer
type Router struct {
	Commands   []*Command
	Prefix     string
	WadlDB     *db.WadlDB
	PermSystem *permissions.PermissionSystem
	Config     *cfg.Config
}

//BuildRouter returns a fully built router stuct with commands preregistered
func BuildRouter(wdb *db.WadlDB, permSystem *permissions.PermissionSystem, cfg *cfg.Config) Router {
	r := Router{
		Prefix:     cfg.Wadl.Prefix,
		WadlDB:     wdb,
		PermSystem: permSystem,
		Config:     cfg,
	}

	return r
}

func (r *Router) SetupPermissions() {
	r.generatePermissionNodes()

	r.PermSystem.AddReferences()
}

//RegisterCommands adds a command(s) to the Router
func (r *Router) RegisterCommands(cmds ...*Command) {
	sort.Slice(cmds, func(i, j int) bool { return cmds[i].Name < cmds[j].Name })
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
			deepestCmd, args, node := findDeepestCommand(cmd, split, cmd.Name)

			ctx := buildContext(r, session, message, deepestCmd, args)

			if !r.userHasCorrectPermissions(session, message.Author, node) {
				ctx.ReplyStringf("You don't have the required permission node `%s` for this cmd.", node)
				return
			}

			deepestCmd.Handler(&ctx)

			//Update UserActivity entry's CommandCount
			var ua model.UserActivity
			tx := db.Instance.DB.Where(&model.UserActivity{UserID: message.Author.ID}).FirstOrCreate(&ua)
			if util.DebugError(tx.Error) {
				log.Error().Err(tx.Error).Msg("An error occured.")
			}
			ua.CommandCount++
			db.Instance.DB.Save(&ua)
		}
	}
}

func (r *Router) userHasCorrectPermissions(session *discordgo.Session, user *discordgo.User, nodeIdentifier string) bool {
	gm, err := session.GuildMember(r.Config.Wadl.GuildID, user.ID)
	if util.DebugError(err) {
		log.Error().Err(err).Msg("An error has occurred.")
		return false
	}

	if r.userHasBypassPermissions(user) {
		log.Debug().Msgf("User %s is on the permission bypass list.", user.ID)
		return true
	}

	return r.PermSystem.UserHasPermissionNode(gm, nodeIdentifier)
}

func (r *Router) userHasBypassPermissions(user *discordgo.User) bool {
	return util.SliceContains(r.Config.Permissions.UserOverride, user.ID)
}

//Finds and returns the deepest subcommand for a given command and arg slice
func findDeepestCommand(prevCmd *Command, args []string, node string) (*Command, []string, string) {
	if len(prevCmd.SubCommands) > 0 {
		if len(args) > 1 {
			found, cmd := triggerCheck(args[1], prevCmd.SubCommands)
			if found {
				return findDeepestCommand(cmd, args[1:], node+"."+cmd.Name)
			}
		}
	}
	return prevCmd, args[1:], node
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

func buildContext(router *Router, session *discordgo.Session, message *discordgo.MessageCreate, command *Command, args []string) Context {
	return *&Context{
		Router:  router,
		Session: session,
		Message: message,
		Command: command,
		Args:    args,
	}
}

func (r *Router) generatePermissionNodes() {
	for _, cmd := range r.Commands {
		cmd.GeneratePermissionNode(r.PermSystem, "")
	}
}

func handlePing(session *discordgo.Session, message *discordgo.MessageCreate) {}
