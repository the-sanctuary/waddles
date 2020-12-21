package waddles

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/internal/commands"
	"github.com/the-sanctuary/waddles/internal/handlers"
	"github.com/the-sanctuary/waddles/pkg/cfg"
	"github.com/the-sanctuary/waddles/pkg/command"
	"github.com/the-sanctuary/waddles/pkg/db"
	"github.com/the-sanctuary/waddles/pkg/handler"
	"github.com/the-sanctuary/waddles/pkg/permissions"
	"github.com/the-sanctuary/waddles/pkg/util"
)

//Waddles .
type Waddles struct {
	//Global Config
	Config   *cfg.Config
	Router   *command.Router
	Database *db.WadlDB
	Session  *discordgo.Session
}

//Run reads the Config, initializes all needed systems, opens the discord api session, and registers the command router and other handlers.
func (w *Waddles) Run() {
	util.InitializeLogging()
	w.Config = cfg.ReadConfig()
	util.SetupLogging()

	var err error

	// Create a Discord session using our bot token (client secret)
	w.Session, err = discordgo.New("Bot " + w.Config.Wadl.Token)
	if util.DebugError(err) {
		log.Fatal().Err(err).Msg("Unable to create a Discord session.  Quitting....")
	}

	// Open connection to database
	w.Database = db.BuildWadlDB(w.Config)
	w.Database.Migrate()

	permSystem := permissions.BuildPermissionSystem(w.Config.GetConfigFileLocation("permissions.toml"))

	r := command.BuildRouter(&w.Database, &permSystem, w.Config)

	r.RegisterCommands(
		commands.Ping,
		commands.Purge,
		commands.Uptime,
		commands.Nitro,
		commands.Debug,
	)

	r.SetupPermissions()

	w.Router = &r

	// Register handlers
	w.Session.AddHandler(w.Router.Handler())
	w.Session.AddHandler(handler.TraceAllMessages)

	w.Session.AddHandler(handlers.UserActivityTextChannel)
	w.Session.AddHandler(handlers.UserActivityVoiceChannel)

	// Open a websocket connection to Discord and start listening
	err = w.Session.Open()

	if util.DebugError(err) {
		log.Fatal().Err(err).Msg("Unable to open a connection to Discord.  Quitting....")
	}

	defer w.Session.Close()

	// Print msg that the bot is running
	log.Info().Msg("Waddles is now running.  Press CTRL-C to quit.")
	util.MarkStartTime()
	util.RegisterTermSignals()
}
