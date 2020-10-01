package main

import (
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
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

	// Open connection to database
	_ = db.NewWadlDB()

	// Register handlers
	session.AddHandler(debugAllMessages)

	// Print msg that the bot is running
	log.Info().Msg("[WADL] Waddles is now running.  Press CTRL-C to quit.")

	util.RegisterTermSignals()
}

func debugAllMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Trace().Msgf("Message Recieved: %s", m.Message.Content)
}
