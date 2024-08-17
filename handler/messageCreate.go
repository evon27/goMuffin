package handler

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// MessageCreate is handlers of messageCreate event
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID && m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, "머핀아 ") {
		content := strings.TrimPrefix(m.Content, "머핀아 ")
		if content == "안녕" {
			s.ChannelMessageSend(m.ChannelID, "안녕")
		}
	} else {
		return
	}
}
