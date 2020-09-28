package handlers

import (
	"github.com/bwmarrin/discordgo"
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
