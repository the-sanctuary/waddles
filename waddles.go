package waddles

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/command"
	"github.com/the-sanctuary/waddles/db"
	"github.com/the-sanctuary/waddles/handler"
	"github.com/the-sanctuary/waddles/permissions"
	"github.com/the-sanctuary/waddles/util"
)

//Run reads the config, initializes all needed systems, opens the discord api session, and registers the command router and other handlers.
func Run() {
	util.InitializeLogging()
	config := util.ReadConfig()
	util.SetupLogging()

	// Create a Discord session using our bot token (client secret)
	session, err := discordgo.New("Bot " + config.Wadl.Token)
	if util.DebugError(err) {
		log.Fatal().Err(err).Msg("Unable to create a Discord session.  Quitting....")
	}

	// Open connection to database
	wdb := db.BuildWadlDB(config)
	wdb.Migrate()

	permSystem := permissions.BuildPermissionSystem(config.GetConfigFileLocation("permissions.toml"))

	router := command.BuildRouter(&wdb, &permSystem, config)

	// Register handlers
	session.AddHandler(router.Handler())
	session.AddHandler(handler.TraceAllMessages)
	session.AddHandler(handler.UserActivityTextChannel)
	session.AddHandler(handler.UserActivityVoiceChannel)

	// Open a websocket connection to Discord and start listening
	err = session.Open()

	if util.DebugError(err) {
		log.Fatal().Err(err).Msg("Unable to open a connection to Discord.  Quitting....")
	}

	defer session.Close()

	// Print msg that the bot is running
	log.Info().Msg("Waddles is now running.  Press CTRL-C to quit.")
	util.MarkStartTime()
	util.RegisterTermSignals()
}
