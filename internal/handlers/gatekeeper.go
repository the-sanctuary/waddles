package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/pkg/cfg"
)

// GatekeeperJoinHandler tracks when a user joins the guild via the GuildMemberAdd event
func GatekeeperJoinHandler(s *discordgo.Session, gma *discordgo.GuildMemberAdd) {
	log.Trace().Msgf("GuildMemberAddEvent - user: %s#%s", gma.Member.User.Username, gma.Member.User.Discriminator)
	// TODO: Add the Newbie role to the newly joined user
	config := cfg.ReadConfig()
	if config.Gatekeeper.Role != "" {
		s.GuildMemberRoleAdd(gma.GuildID, gma.User.ID, config.Gatekeeper.Role)
	}
}

// GatekeeperMsgHandler tracks accepting or declining of server rules in the specified gatekeeper channel
// func GatekeeperMsgHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	config := cfg.ReadConfig()
// 	log.Trace().Msgf(config.Gatekeeper.Channel)

// 	// Only do this if someone typed in the gatekeeper channel
// 	if m.ChannelID == config.Gatekeeper.Channel {
// 		log.Trace().Msgf("Do we get here? [0]")

// 		// Always delete messages from users in the channel before exiting the function
// 		defer s.ChannelMessageDelete(m.ChannelID, m.ID)

// 		if m.Content == "accept" {
// 			log.Trace().Msgf("User %s#%s accepted the rules.", m.Author.Username, m.Author.Discriminator)
// 			// TODO: Remove the Newbie role
// 		} else {
// 			log.Trace().Msgf("User %s#%s declined the rules.", m.Author.Username, m.Author.Discriminator)
// 			// TODO: Kick the user from the server
// 			return
// 		}
// 	}

// }
