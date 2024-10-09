package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
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
		if content == "안녕" {
			s.ChannelMessageSend(m.ChannelID, "안녕")
		}
	} else {
		return
	}
}
