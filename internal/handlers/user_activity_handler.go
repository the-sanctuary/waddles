package handlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"

	"github.com/the-sanctuary/waddles/pkg/db"
	"github.com/the-sanctuary/waddles/pkg/util"
)

//UserActivityTextChannel tracks when a user sends a message to a text channel
func UserActivityTextChannel(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}
	var ua db.UserActivity

	r := db.Instance.FirstOrInit(&ua, db.UserActivity{User: db.User{DiscordID: m.Author.ID}})
	util.DebugError(r.Error)

	now := time.Now()
	ua.LastChannelTextAppearence = &now
	ua.MessageCount++

	db.Instance.Save(&ua)
}

//UserActivityVoiceChannel tracks when users join or leave a voice channel
func UserActivityVoiceChannel(s *discordgo.Session, vsu *discordgo.VoiceStateUpdate) {
	if vsu.UserID == s.State.User.ID {
		return
	}
	var ua db.UserActivity

	log.Debug().Msgf("Voice session id: %s", vsu.SessionID)

	r := db.Instance.FirstOrInit(&ua, db.UserActivity{User: db.User{DiscordID: vsu.UserID}})
	util.DebugError(r.Error)

	now := time.Now()
	ua.LastChannelVoiceAppearence = &now

	_, err := s.State.VoiceState(vsu.GuildID, vsu.UserID)
	if err == discordgo.ErrStateNotFound {
		ua.VoiceCount++
		log.Trace().Msgf("UserVoiceAcitivtyUpdated - user: %s", vsu.UserID)
	} else if err != nil {
		util.DebugError(err)
	}

	db.Instance.Save(&ua)
}
