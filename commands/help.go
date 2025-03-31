package commands

import (
	"strings"

	"git.wh64.net/muffin/goMuffin/configs"
	"git.wh64.net/muffin/goMuffin/utils"
	"github.com/bwmarrin/discordgo"
)

var HelpCommand *Command = &Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "도움말",
		Description: "기본적인 사용ㅂ법이에요.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "명령어",
				Description: "해당 명령어에 대ㅎ한 도움말을 볼 수 있어요.",
				Choices: func() []*discordgo.ApplicationCommandOptionChoice {
					choices := []*discordgo.ApplicationCommandOptionChoice{}
					for _, command := range Discommand.Commands {
						choices = append(choices, &discordgo.ApplicationCommandOptionChoice{
							Name:  command.Name,
							Value: command.Name,
						})
					}
					return choices
				}(),
			},
		},
	},
	Aliases: []string{"도움", "명령어", "help"},
	DetailedDescription: &DetailedDescription{
		Usage:    "머핀아 도움말 [명령어]",
		Examples: []string{"머핀아 도움말", "머핀아 도움말 배워"},
	},
	Category: Generals,
	MessageRun: func(ctx *MsgContext) {
		helpRun(ctx.Command, ctx.Session, ctx.Msg, &ctx.Args)
	},
	ChatInputRun: func(ctx *ChatInputContext) {
		var args *[]string
		helpRun(ctx.Command, ctx.Session, ctx.Inter, args)
	},
}

func getCommandsByCategory(d *DiscommandStruct, category Category) []string {
	commands := []string{}
	for _, command := range d.Commands {
		if command.Category == category {
			commands = append(commands, "- "+command.Name+": "+command.Description)
		}
	}
	return commands
}

func helpRun(c *Command, s *discordgo.Session, m any, args *[]string) {
	var commandName string
	embed := &discordgo.MessageEmbed{
		Color: int(utils.EDefault),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "버전:" + configs.MUFFIN_VERSION,
		},
		Thumbnail: &discordgo.MessageEmbedThumbnail{
			URL: s.State.User.AvatarURL("512"),
		},
	}

	switch m := m.(type) {
	case *discordgo.MessageCreate:
		commandName = Discommand.Aliases[strings.Join(*args, " ")]
	case *discordgo.InteractionCreate:
		optsMap := map[string]*discordgo.ApplicationCommandInteractionDataOption{}
		for _, opt := range m.ApplicationCommandData().Options {
			optsMap[opt.Name] = opt
		}
		if opt, ok := optsMap["명령어"]; ok {
			commandName = opt.StringValue()
		}
	}

	if commandName == "" || Discommand.Commands[commandName] == nil {
		embed.Title = s.State.User.Username + "의 도움말"
		embed.Description = utils.CodeBlockWithLanguage(
			"md",
			"# 일반\n"+
				strings.Join(getCommandsByCategory(Discommand, Generals), "\n")+
				"\n\n# 채팅\n"+
				strings.Join(getCommandsByCategory(Discommand, Chattings), "\n"),
		)

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
		return
	}

	command := Discommand.Commands[commandName]

	embed.Title = s.State.User.Username + "의 " + command.Name + " 도움말"
	embed.Fields = []*discordgo.MessageEmbedField{
		{
			Name:   "설명",
			Value:  utils.InlineCode(command.Description),
			Inline: true,
		},
		{
			Name:   "사용법",
			Value:  utils.InlineCode(command.DetailedDescription.Usage),
			Inline: true,
		},
	}

	if command.Aliases != nil {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "별칭",
			Value: utils.CodeBlockWithLanguage("md", strings.Join(addPrefix(command.Aliases), "\n")),
		})
	} else {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "별칭",
			Value: "없음",
		})
	}

	if command.DetailedDescription.Examples != nil {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "예시",
			Value: utils.CodeBlockWithLanguage("md", strings.Join(addPrefix(c.DetailedDescription.Examples), "\n")),
		})
	} else {
		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:  "예시",
			Value: "없음",
		})
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
