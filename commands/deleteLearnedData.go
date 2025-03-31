package commands

import (
	"context"
	"strconv"
	"strings"

	"git.wh64.net/muffin/goMuffin/databases"
	"git.wh64.net/muffin/goMuffin/utils"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var DeleteLearnedDataCommand *Command = &Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "삭제",
		Description: "당신이 가르쳐준 단ㅇ어를 삭제해요.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Name:        "단어",
				Description: "삭제할 단어를 입ㄹ력해주세요.",
				Required:    true,
			},
		},
	},
	Aliases: []string{"잊어", "지워"},
	DetailedDescription: &DetailedDescription{
		Usage:    "머핀아 삭제 (삭제할 단어)",
		Examples: []string{"머핀아 삭제 머핀"},
	},
	Category: Chattings,
	MessageRun: func(ctx *MsgContext) {
		deleteLearnedDataRun(ctx.Command, ctx.Session, ctx.Msg, &ctx.Args)
	},
	ChatInputRun: func(ctx *ChatInputContext) {
		var args *[]string
		deleteLearnedDataRun(ctx.Command, ctx.Session, ctx.Inter, args)
	},
}

func deleteLearnedDataRun(c *Command, s *discordgo.Session, m any, args *[]string) {
	var command, userId, description string
	var datas []databases.Learn
	var options []discordgo.SelectMenuOption

	switch m := m.(type) {
	case *discordgo.MessageCreate:
		command = strings.Join(*args, " ")
		userId = m.Author.ID

		if command == "" {
			s.ChannelMessageSendEmbedReply(m.ChannelID, &discordgo.MessageEmbed{
				Title:       "❌ 오류",
				Description: "올바르지 않ㅇ은 용법이에요.",
				Fields: []*discordgo.MessageEmbedField{
					{
						Name:  "사용법",
						Value: utils.InlineCode(c.DetailedDescription.Usage),
					},
					{
						Name:  "예시",
						Value: utils.CodeBlockWithLanguage("md", strings.Join(addPrefix(c.DetailedDescription.Examples), "\n")),
					},
				},
				Color: int(utils.EFail),
			}, m.Reference())
		}
	case *discordgo.InteractionCreate:
		s.InteractionRespond(m.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Flags: discordgo.MessageFlagsEphemeral,
			},
		})

		optsMap := map[string]*discordgo.ApplicationCommandInteractionDataOption{}
		for _, opt := range m.ApplicationCommandData().Options {
			optsMap[opt.Name] = opt
		}
		if opt, ok := optsMap["단어"]; ok {
			command = opt.StringValue()
		}
		userId = m.Member.User.ID
	}

	cur, err := databases.Learns.Find(context.TODO(), bson.M{"user_id": userId, "command": command})
	if err != nil {
		embed := &discordgo.MessageEmbed{
			Title: "❌ 오류",
			Color: int(utils.EFail),
		}
		if err == mongo.ErrNoDocuments {
			embed.Description = "해당 하는 지식ㅇ을 찾을 수 없어요."
			switch m := m.(type) {
			case *discordgo.MessageCreate:
				s.ChannelMessageSendEmbedReply(m.ChannelID, embed, m.Reference())
			case *discordgo.InteractionCreate:
				s.InteractionResponseEdit(m.Interaction, &discordgo.WebhookEdit{
					Embeds: &[]*discordgo.MessageEmbed{embed},
				})
			}
			return
		}

		embed.Description = "데이터를 가져오는데 실패했어요."
		switch m := m.(type) {
		case *discordgo.MessageCreate:
			s.ChannelMessageSendEmbedReply(m.ChannelID, embed, m.Reference())
		case *discordgo.InteractionCreate:
			s.InteractionResponseEdit(m.Interaction, &discordgo.WebhookEdit{
				Embeds: &[]*discordgo.MessageEmbed{embed},
			})
		}
		return
	}

	cur.All(context.TODO(), &datas)

	for i := range len(datas) {
		data := datas[i]

		options = append(options, discordgo.SelectMenuOption{
			Label:       strconv.Itoa(i+1) + "번 지식",
			Description: data.Result,
			Value:       utils.DeleteLearnedData + data.Id.Hex() + `&No.` + strconv.Itoa(i+1),
		})
		description += strconv.Itoa(i+1) + ". " + data.Result + "\n"
	}

	embed := &discordgo.MessageEmbed{
		Title: command + " 삭제",
		Description: utils.CodeBlockWithLanguage("md", "# "+command+" 에 대한 대답 중 하나를 선ㅌ택하여 삭제해주세요.\n"+
			description),
		Color: int(utils.EDefault),
	}

	components := []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.SelectMenu{
					MenuType:    discordgo.StringSelectMenu,
					CustomID:    utils.DeleteLearnedDataUserId + userId,
					Options:     options,
					Placeholder: "ㅈ지울 응답을 선택해주세요.",
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					CustomID: utils.DeleteLearnedDataCancel + userId,
					Label:    "취소하기",
					Style:    discordgo.DangerButton,
					Disabled: false,
				},
			},
		},
	}

	switch m := m.(type) {
	case *discordgo.MessageCreate:
		s.ChannelMessageSendComplex(m.ChannelID, &discordgo.MessageSend{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: components,
			Reference:  m.Reference(),
		})
	case *discordgo.InteractionCreate:
		s.InteractionResponseEdit(m.Interaction, &discordgo.WebhookEdit{
			Embeds:     &[]*discordgo.MessageEmbed{embed},
			Components: &components,
		})
	}
}
