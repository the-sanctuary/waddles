package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/command"
	"github.com/the-sanctuary/waddles/handler"
	"github.com/the-sanctuary/waddles/db"
	"github.com/the-sanctuary/waddles/util"
)

func main() {
	util.InitializeLogging()
	util.ReadConfig()
	util.SetupLogging()

	// Create a Discord session using our bot token (client secret)
	session, err := discordgo.New("Bot " + util.Cfg.Wadl.Token)
	if util.DebugError(err) {
		log.Info().Msg("[WADL] Unable to create a Discord session.  Quitting....")
		log.Debug().Msg("[IERR] " + err.Error())
		os.Exit(1)
	}

	// Open a websocket connection to Discord and start listening
	err = session.Open()
	if util.DebugError(err) {
		log.Info().Msg("[WADL] Unable to open a connection to Discord.  Quitting....")
		os.Exit(1)
	}
	defer session.Close()

	router := command.BuildRouter()
	// Open connection to database
	_ = db.NewWadlDB()

	// Register handlers
	session.AddHandler(handler.TraceAllMessages)
	session.AddHandler(router.Handler())

	// Print msg that the bot is running
	log.Info().Msg("[WADL] Waddles is now running.  Press CTRL-C to quit.")

	util.RegisterTermSignals()
}
