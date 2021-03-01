package handlers

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

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

//NicknameUpdateListener logs every member's nickname  changes in the database.
func NicknameUpdateListener(s *discordgo.Session, gmu *discordgo.GuildMemberUpdate) {
	var nnu db.NicknameUpdate

	tx := db.Instance.Order("created_at DESC").Where("discord_id = ?, gmu.User.ID").First(&nnu)

	if util.DebugError(tx.Error) && tx.Error == gorm.ErrRecordNotFound {
		tx = db.Instance.Create(&db.NicknameUpdate{Nickname: gmu.Nick, User: db.User{DiscordID: gmu.User.ID}})
		log.Trace().Msgf("Started tracking nickname updates for %s (%s#%s)", gmu.User.ID, gmu.User.Username, gmu.User.Discriminator)
		util.DebugError(tx.Error)
	}

	log.Trace().Msgf("Filed nickname update for %s (%s -> %s)")

}
