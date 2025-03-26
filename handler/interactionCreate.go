package handler

import (
	"git.wh64.net/muffin/goMuffin/commands"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	commands.Discommand.ChatInputRun(i.ApplicationCommandData().Name, s, i)
	return
}
