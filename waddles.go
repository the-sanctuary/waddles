package waddles

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/command"
	"github.com/the-sanctuary/waddles/db"
	"github.com/the-sanctuary/waddles/handler"
	"github.com/the-sanctuary/waddles/permissions"
	"github.com/the-sanctuary/waddles/util"
)

//Waddles holds the state for a Waddles session
type Waddles struct {
	WadlDB  *db.WadlDB
	Session *discordgo.Session
	Router  *command.Router
}

//Run reads the config, initializes all needed systems, opens the discord api session, and registers the command router and other handlers.
func Run() {
	w := Waddles{}

	util.InitializeLogging()
	util.ReadConfig()
	util.SetupLogging()

	// Create a Discord session using our bot token (client secret)
	var err error
	w.Session, err = discordgo.New("Bot " + util.Cfg.Wadl.Token)
	if util.DebugError(err) {
		log.Info().Msg("[WADL] Unable to create a Discord session.  Quitting....")
		os.Exit(1)
	}

	// Open connection to database
	wdb := db.BuildWadlDB()
	wdb.Migrate()

	db.Instance = &wdb
	w.WadlDB = &wdb

	permSystem := permissions.BuildPermissionSystem("./permissions.toml") //TODO: remove hard coded location to permissions.toml

	router := command.BuildRouter(w.WadlDB, permSystem)

	w.Router = &router

	// Register handlers
	w.Session.AddHandler(router.Handler())
	w.Session.AddHandler(handler.TraceAllMessages)
	w.Session.AddHandler(handler.UserActivityTextChannel)
	w.Session.AddHandler(handler.UserActivityVoiceChannel)

	// Open a websocket connection to Discord and start listening
	err = w.Session.Open()
	defer w.Session.Close()
	if util.DebugError(err) {
		log.Info().Msg("[WADL] Unable to open a connection to Discord.  Quitting....")
		os.Exit(1)
	}

	// Print msg that the bot is running
	log.Info().Msg("[WADL] Waddles is now running.  Press CTRL-C to quit.")
	util.MarkStartTime()
	util.RegisterTermSignals()
}
