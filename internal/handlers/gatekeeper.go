package handlers

import (
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/pkg/cfg"
	"github.com/the-sanctuary/waddles/pkg/util"
)

// GatekeeperJoinHandler tracks when a user joins the guild via the GuildMemberAdd event
func GatekeeperJoinHandler(s *discordgo.Session, gma *discordgo.GuildMemberAdd) {
	log.Trace().Msgf("GuildMemberAddEvent - user: %s#%s", gma.Member.User.Username, gma.Member.User.Discriminator)

	if cfg.Cfg().Gatekeeper.RoleID != "" {
		s.GuildMemberRoleAdd(gma.GuildID, gma.User.ID, cfg.Cfg().Gatekeeper.RoleID)
	}
}

//GatekeeperMsgHandler tracks accepting or declining of server rules in the specified gatekeeper channel
func GatekeeperMsgHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Only do this if someone typed in the gatekeeper channel
	if m.ChannelID == cfg.Cfg().Gatekeeper.ChannelID {

		// Always delete messages from users in the channel before exiting the function
		defer s.ChannelMessageDelete(m.ChannelID, m.ID)

		if strings.ToLower(m.Content) == "accept" {
			log.Trace().Msgf("User %s#%s accepted the rules.", m.Author.Username, m.Author.Discriminator)
			s.GuildMemberRoleRemove(cfg.Cfg().Wadl.GuildID, m.Author.ID, cfg.Cfg().Gatekeeper.RoleID)
		} else if strings.ToLower(m.Content) == "decline" {
			log.Trace().Msgf("User %s#%s declined the rules.", m.Author.Username, m.Author.Discriminator)
			s.GuildMemberDelete(cfg.Cfg().Wadl.GuildID, m.Author.ID)
			return
		} else {
			errMsg, err := s.ChannelMessageSend(m.ChannelID, "Please enter accept/decline.")

			time.Sleep(time.Second * 3)
			s.ChannelMessageDelete(m.ChannelID, errMsg.ID)

			if util.DebugError(err) {
				return
			}
		}
	}
}
