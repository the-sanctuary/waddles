package handlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/the-sanctuary/waddles/pkg/db"
	"github.com/the-sanctuary/waddles/pkg/util"
)

//UserActivityTextChannel
func UserActivityTextChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	var ua db.UserActivity

	r := db.Instance.Where("discord_id = ?", m.Author.ID).FirstOrCreate(&ua)
	util.DebugError(r.Error)

	now := time.Now()
	ua.LastChannelTextAppearence = &now
	ua.MessageCount++

	db.Instance.Save(&ua)
}

//UserActivityTextChannel
func UserActivityVoiceChannel(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
	if vsu.UserID == s.State.User.ID {
		return
	}
	var ua db.UserActivity

	log.Debug().Msgf("Voice session id: %s", vsu.SessionID)

	r := db.Instance.Where("discord_id = ?", vsu.UserID).FirstOrCreate(&ua)
	util.DebugError(r.Error)

	now := time.Now()
	ua.LastChannelVoiceAppearence = &now
	ua.VoiceCount++
	//TODO Figure out how to distinguish join/leave events without having a flag stored in the DB

	db.Instance.Save(&ua)
}
