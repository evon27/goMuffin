package commands

import (
	"runtime"

	"git.wh64.net/muffin/goMuffin/configs"
	"git.wh64.net/muffin/goMuffin/utils"
	"github.com/bwmarrin/discordgo"
)

var InformationCommand *Command = &Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "정보",
		Description: "머핀봇의 정보를 알ㄹ려줘요.",
	},
	DetailedDescription: &DetailedDescription{
		Usage: "머핀아 정보",
	},
	Category: Generals,
	MessageRun: func(ctx *MsgContext) {
		informationRun(ctx.Session, ctx.Msg)
	},
	ChatInputRun: func(ctx *ChatInputContext) {
		informationRun(ctx.Session, ctx.Inter)
	},
}

func informationRun(s *discordgo.Session, m any) {
	owner, _ := s.User(configs.Config.Bot.OwnerId)
	embed := &discordgo.MessageEmbed{
		Title: s.State.User.Username + "의 정보",
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:  "운영 체제",
				Value: utils.InlineCode(runtime.GOOS + " " + runtime.GOARCH),
			},
			{
				Name:  "제작자",
				Value: owner.Username,
			},
			{
				Name:  "버전",
				Value: utils.InlineCode(configs.MUFFIN_VERSION),
			},
			{
				Name:   "최근에 업데이트된 날짜",
				Value:  utils.TimeWithStyle(configs.UpdatedAt, utils.RelativeTime),
				Inline: true,
			},
			{
				Name:   "업타임",
				Value:  utils.TimeWithStyle(configs.StartedAt, utils.RelativeTime),
				Inline: true,
			},
		},
		Color: int(utils.EDefault),
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: s.State.User.AvatarURL("512"),
		},
	}

	switch m := m.(type) {
	case *discordgo.MessageCreate:
		s.ChannelMessageSendEmbedReply(m.ChannelID, embed, m.Reference())
	case *discordgo.InteractionCreate:
		s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{embed},
			},
		})
	}
}
