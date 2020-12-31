package commands

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/the-sanctuary/waddles/internal/model"
	"github.com/the-sanctuary/waddles/pkg/cmd"

	"github.com/the-sanctuary/waddles/pkg/util"
)

var NitroChannel *cmd.Command = &cmd.Command{
	Name:        "channel",
	Aliases:     *&[]string{"c"},
	Description: "control your voice channel",
	Usage:       "channel (register|release)",
	SubCommands: []*cmd.Command{NitroChannelRegister, NitroChannelRelease},
	Handler: func(c *cmd.Context) {
		//Check to see if a user already has a channel registered
		var chann model.NitroUserChannel
		c.DB().Where(&model.NitroUserChannel{UserID: c.Message.Author.ID}).First(&chann)

		if chann.UserID == "" {
			c.ReplyString("You don't have a channel.")
		} else {
			c.ReplyStringf("Your channel is named: %s", chann.Name)
		}
	},
}

var NitroChannelRelease *cmd.Command = &cmd.Command{
	Name:        "release",
	Aliases:     *&[]string{"rl"},
	Description: "release your voice channel",
	Usage:       "release",
	Handler: func(c *cmd.Context) {
		//Check to see if a user already has a channel registered
		var chann model.NitroUserChannel
		c.DB().Where(&model.NitroUserChannel{UserID: c.Message.Author.ID}).First(&chann)

		if chann.UserID == "" {
			c.ReplyString("You don't have a channel to release!")
		} else {
			c.Session.ChannelDelete(chann.ChannelID)

			c.DB().Delete(&chann).Commit()

			c.ReplyString("Your channel has been released.")
		}
	},
}

var NitroChannelRegister *cmd.Command = &cmd.Command{
	Name:        "register",
	Aliases:     *&[]string{"r"},
	Description: "register your voice channel",
	Usage:       "register <name>",
	Handler: func(c *cmd.Context) {
		if len(c.Args) < 1 {
			c.ReplyString("You must name supply a name for your channel")
			return
		}

		//Check to see if a user already has a channel registered
		var chann model.NitroUserChannel
		c.DB().Where(&model.NitroUserChannel{UserID: c.Message.Author.ID}).First(&chann)

		if chann.UserID == "" {
			channelName := strings.Join(c.Args, " ")

			if len(channelName) > 100 || len(channelName) < 4 {
				c.ReplyStringf("Channel Name: `%s` is invalid. The length is out of bounds (4 < name < 100", channelName)
				return
			}

			permOverwrite := discordgo.PermissionOverwrite{
				ID:    c.Message.Author.ID,
				Type:  "1",
				Allow: discordgo.PermissionManageChannels,
			}

			createdChannel, err := c.Session.GuildChannelCreateComplex(c.Message.GuildID, discordgo.GuildChannelCreateData{
				Name:                 channelName,
				Type:                 discordgo.ChannelTypeGuildVoice,
				ParentID:             c.Router.Config.NitroPerk.BoosterChannel.ParentID,
				PermissionOverwrites: []*discordgo.PermissionOverwrite{&permOverwrite},
			})

			if err != nil {
				c.ReplyString("An error occured while trying to make your channel. Please try again. If this issue persists, contact an admin.")
				util.DebugError(err)
				return
			}

			// Make sure that the nitrochannel object is initialized properly.
			chann.UserID = c.Message.Author.ID
			chann.ChannelID = createdChannel.ID
			chann.Name = createdChannel.Name
			chann.Active = true

			// Add the server to the database
			c.DB().Create(&chann)

			c.ReplyStringf("Your channel: `%s` has been registered.", channelName)
		} else {
			c.ReplyString("Sorry, you can only have one channel at a time. Please use the `release` subcommand to release your previous channel before `regsiter`ing a new one")
		}
	},
}
