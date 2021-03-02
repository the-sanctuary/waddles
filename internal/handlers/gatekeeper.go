package handlers

import (
	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog/log"
)

// GatekeeperHandler tracks when a user joins the guild via the GuildMemberAdd event
func GatekeeperHandler(s *discordgo.Session, gma *discordgo.GuildMemberAdd) {
	log.Trace().Msgf("GuildMemberAddEvent - user: %s#%s", gma.Member.User.Username, gma.Member.User.Discriminator)
}
