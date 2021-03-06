package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

// TraceAllMessages sends a Trace message to the log with showing every message processed
func TraceAllMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Ignore all messages from bots
	if m.Author.Bot {
		return
	}

	log.Trace().Msgf("Message Recieved: %s", m.Message.Content)
}
