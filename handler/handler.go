package handler

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

func MsgHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all message if they come from Waddles
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
}

// TraceAllMessages sends a Trace message to the log with showing every message processed
func TraceAllMessages(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	log.Trace().Msgf("Message Recieved: %s", m.Message.Content)
}
