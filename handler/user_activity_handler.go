package handler

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/db"
	"github.com/the-sanctuary/waddles/util"
)

func UserActivityTextChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	var ua db.UserActivity

	r := db.CurrentWadlDB().DB.Where(&db.UserActivity{UserID: m.Author.ID}).FirstOrCreate(&ua)
	util.DebugError(r.Error)

	now := time.Now()
	ua.LastChannelTextAppearence = &now
	ua.MessageCount++

	db.CurrentWadlDB().DB.Save(&ua)
}

//TODO Figure out how to distinguish join/leave events without having a flag stored in the DB
func UserActivityVoiceChannel(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
	if vsu.UserID == s.State.User.ID {
		return
	}
	var ua db.UserActivity

	log.Debug().Msgf("Voice session id: %s", vsu.SessionID)

	r := db.CurrentWadlDB().DB.Where(&db.UserActivity{UserID: vsu.UserID}).FirstOrCreate(&ua)
	util.DebugError(r.Error)

	now := time.Now()
	ua.LastChannelVoiceAppearence = &now
	ua.VoiceCount++

	db.CurrentWadlDB().DB.Save(&ua)
}
