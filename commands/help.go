package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var HelpCommand Command = Command{
	Name:        "도움말",
	Aliases:     []string{"도움", "명령어", "help"},
	Description: "기본적인 사용ㅂ법이에요.",
	DetailedDescription: DetailedDescription{
		Usage:    "머핀아 도움말 [명령어]",
		Examples: []string{"머핀아 도움말", "머핀아 도움말 배워"},
	},
}

func (c *Command) helpMessageRun(s *discordgo.Session, m *discordgo.MessageCreate) {
	fmt.Println(c.Name)
	s.ChannelMessageSend(m.ChannelID, "asdf")
}
