package handler

import (
	"time"

	"github.com/bwmarrin/discordgo"
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

func UserActivityVoiceChannel(s *discordgo.Session, m *discordgo.VoiceStateUpdate) {
	if m.UserID == s.State.User.ID {
		return
	}
	var ua db.UserActivity

	db.CurrentWadlDB().DB.Where(&db.UserActivity{UserID: m.UserID}).FirstOrCreate(&ua)

	now := time.Now()
	ua.LastChannelVoiceAppearence = &now

	db.CurrentWadlDB().DB.Save(&ua)
}
