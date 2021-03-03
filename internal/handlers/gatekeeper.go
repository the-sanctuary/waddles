package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

// GatekeeperJoinHandler tracks when a user joins the guild via the GuildMemberAdd event
func GatekeeperJoinHandler(s *discordgo.Session, gma *discordgo.GuildMemberAdd) {
	log.Trace().Msgf("GuildMemberAddEvent - user: %s#%s", gma.Member.User.Username, gma.Member.User.Discriminator)
}

// GatekeeperMsgHandler tracks accepting or declining of server rules in the specified gatekeeper channel
// func GatekeeperMsgHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
// 	config := cfg.ReadConfig()
// 	log.Trace().Msgf("Are we making it here?!")

// 	if config.Gatekeeper.Channel != m.ChannelID {
// 		log.Trace().Msgf("Are we making it here now?!")
// 		return
// 	}

// 	if m.Message.Content == "accept" {
// 		log.Trace().Msgf("User %s#%s accepted the rules", m.Author.Username, m.Author.Discriminator)
// 		// TODO: Add removal of newbie role code here
// 	} else {
// 		log.Trace().Msgf("User %s#%s declined the rules", m.Author.Username, m.Author.Discriminator)
// 		// TODO: Add code to kick user here
// 	}
// 	//s.ChannelMessageDelete(m.ChannelID, m.Message.ID)
// }
