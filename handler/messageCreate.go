package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
	"gitlab.wh64.net/muffin/goMuffin/commands"
	"gitlab.wh64.net/muffin/goMuffin/configs"
)

// MessageCreate is handlers of messageCreate event
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	config := configs.Config
	if m.Author.ID == s.State.User.ID && m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, config.Prefix) {
		content := strings.TrimPrefix(m.Content, config.Prefix)
		command := commands.Discommand.Aliases[content]

		if command == "" {
			return
		}

		commands.Discommand.MessageRun(command, s, m)
	} else {
		return
	}
}
