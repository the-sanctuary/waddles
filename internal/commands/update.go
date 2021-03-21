package commands

import (
	"fmt"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/the-sanctuary/waddles/pkg/cfg"
	"github.com/the-sanctuary/waddles/pkg/cmd"
	"github.com/the-sanctuary/waddles/pkg/util"
)

var update *cmd.Command = &cmd.Command{
	Name:        "update",
	Aliases:     []string{"u"},
	Description: "Force the bot to perform various update functions",
	Usage:       "update [subsystem]",
	SubCommands: []*cmd.Command{updateGatekeeper, updateConfig},
}

var updateConfig *cmd.Command = &cmd.Command{
	Name:        "config",
	Aliases:     []string{"cfg"},
	Description: "Reload the config file from disk.",
	Usage:       "config",
	Handler: func(c *cmd.Context) {
		log.Info().Msgf("User %s force reloaded the config from disk.", c.Message.Author.ID)
		cfg.ReloadCfgFromDisk()
		c.ReplyTimeDeleteStringf(3*time.Second, "Config reloaded.")
	},
}

var updateGatekeeper *cmd.Command = &cmd.Command{
	Name:        "gatekeeper",
	Aliases:     []string{"gk"},
	Description: "Update gatekeeper settings from the config file",
	Usage:       "gatekeeper",
	Handler: func(c *cmd.Context) {
		/*
		 * Read in the gatekeeper settings from the config file (waddles.toml)
		 * Unfortunately, this reads in the whole config file, so perhaps we should
		 * 	consider separating the gatekeeper stuff out to its own config for this
		 * 	purpose
		 */
		config := cfg.ReadConfig()

		// Write the gatekeeper portion of the config into the config passed from the router
		c.Router.Config.Gatekeeper = config.Gatekeeper

		// Trim any whitespace at the end, as well as any trailing newlines and/or carriage returns
		config.Gatekeeper.Rules = strings.TrimRight(config.Gatekeeper.Rules, " ")
		config.Gatekeeper.Rules = strings.TrimSuffix(config.Gatekeeper.Rules, "\n")

		// Delete any messages in the gatekeeper channel specified by the config file
		// The rate limit (20) here is arbitrary, but it should be at least 1
		messages, err := c.Session.ChannelMessages(config.Gatekeeper.ChannelID, 20, "", "", "")
		if util.DebugError(err) {
			c.ReplyString("An error occured. Check the log for details.")
		}

		var msgIds []string
		for _, m := range messages {
			msgIds = append(msgIds, m.ID)
		}

		if len(msgIds) == 1 {
			c.Session.ChannelMessageDelete(config.Gatekeeper.ChannelID, msgIds[0])
		} else {
			c.Session.ChannelMessagesBulkDelete(config.Gatekeeper.ChannelID, msgIds)
		}

		// Build the message from the gatekeeper config info
		msg := fmt.Sprintf("%s\n```\n%s```\n", config.Gatekeeper.WelcomeMsg, config.Gatekeeper.Rules)
		msg += fmt.Sprintf("By typing accept, you agree to the rules listed here, and will abide by them at all times while in the server.  " +
			"You may decline to accept these rules, but you will be not be granted access to the server, and will instead be kicked.\n\n" +
			"Please type `accept`, to accept the rules above, or `decline`, to leave the server.")

		log.Trace().Msgf("%s", msg)

		// Send the message to the channel specified in the gatekeeper config
		c.Session.ChannelMessageSend(config.Gatekeeper.ChannelID, msg)
	},
}
