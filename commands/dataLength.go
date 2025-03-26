package commands

import (
	"context"
	"log"
	"strconv"

	"git.wh64.net/muffin/goMuffin/databases"
	"git.wh64.net/muffin/goMuffin/utils"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type chStruct struct {
	name   dataType
	length int
}

type dataType int

const (
	text dataType = iota
	muffin
	nsfw
	learn
	userLearn
)

var DataLengthCommand *Command = &Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Type:        discordgo.ChatApplicationCommand,
		Name:        "데이터학습량",
		Description: "봇이 학습한 데ㅇ이터량을 보여줘요.",
	},
	Aliases: []string{"학습데이터량", "데이터량", "학습량"},
	DetailedDescription: &DetailedDescription{
		Usage: "머핀아 학습데이터량",
	},
}
var ch chan chStruct = make(chan chStruct)

func getLength(data dataType, userId string) {
	var err error
	var cur *mongo.Cursor
	var datas []bson.M

	switch data {
	case text:
		cur, err = databases.Texts.Find(context.TODO(), bson.D{{}})
	case muffin:
		cur, err = databases.Texts.Find(context.TODO(), bson.D{{Key: "persona", Value: "muffin"}})
	case nsfw:
		cur, err = databases.Texts.Find(context.TODO(), bson.D{
			{
				Key: "persona",
				Value: bson.M{
					"$regex": "^user",
				},
			},
		})
	case learn:
		cur, err = databases.Learns.Find(context.TODO(), bson.D{{}})
	case userLearn:
		cur, err = databases.Learns.Find(context.TODO(), bson.D{{Key: "user_id", Value: userId}})
	}
	if err != nil {
		log.Fatalln(err)
	}

	defer cur.Close(context.TODO())

	cur.All(context.TODO(), &datas)
	ch <- chStruct{name: data, length: len(datas)}
}

func (c *Command) dataLengthRun(s *discordgo.Session, m interface{}) {
	var i *discordgo.Interaction
	var referance *discordgo.MessageReference
	var username, userId, channelId string
	var textLength,
		muffinLength,
		nsfwLength,
		learnLength,
		userLearnLength int

	switch m := m.(type) {
	case *discordgo.MessageCreate:
		username = m.Author.Username
		userId = m.Author.ID
		channelId = m.ChannelID
		referance = m.Reference()
	case *discordgo.InteractionCreate:
		username = m.Member.User.Username
		userId = m.Member.User.ID
		channelId = m.ChannelID
		i = m.Interaction
		s.InteractionRespond(i,
			&discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseDeferredChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Flags: discordgo.MessageFlagsEphemeral,
				},
			})
	}

	go getLength(text, "")
	go getLength(muffin, "")
	go getLength(nsfw, "")
	go getLength(learn, "")
	go getLength(userLearn, userId)

	for i := 0; i < 5; i++ {
		resp := <-ch
		switch dataType(resp.name) {
		case text:
			textLength = resp.length
		case muffin:
			muffinLength = resp.length
		case nsfw:
			nsfwLength = resp.length
		case learn:
			learnLength = resp.length
		case userLearn:
			userLearnLength = resp.length
		}
	}

	sum := textLength + learnLength

	// 나중에 djs처럼 Embed 만들어 주는 함수 만들어야겠다
	// 지금은 임시방편
	embed := &discordgo.MessageEmbed{
		Title:       "저장된 데이터량",
		Description: "총합: " + utils.InlineCode(strconv.Itoa(sum)) + "개",
		Color:       int(utils.EDefault),
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "총 채팅 데이터량",
				Value:  utils.InlineCode(strconv.Itoa(textLength)) + "개",
				Inline: true,
			},
			{
				Name:   "총 지식 데이터량",
				Value:  utils.InlineCode(strconv.Itoa(learnLength)) + "개",
				Inline: true,
			},
			{
				Name:  "머핀 데이터량",
				Value: utils.InlineCode(strconv.Itoa(muffinLength)) + "개",
			},
			{
				Name:   "nsfw 데이터량",
				Value:  utils.InlineCode(strconv.Itoa(nsfwLength)) + "개",
				Inline: true,
			},
			{
				Name:   username + "님이 가르쳐준 데이터량",
				Value:  utils.InlineCode(strconv.Itoa(userLearnLength)) + "개",
				Inline: true,
			},
		},
	}

	switch m.(type) {
	case *discordgo.MessageCreate:
		s.ChannelMessageSendEmbedReply(channelId, embed, referance)
	case *discordgo.InteractionCreate:
		s.InteractionResponseEdit(i, &discordgo.WebhookEdit{
			Embeds: &[]*discordgo.MessageEmbed{embed},
		})
	}
}

func (c *Command) dataLengthMessageRun(s *discordgo.Session, m *discordgo.MessageCreate) {
	c.dataLengthRun(s, m)
}

func (c *Command) dataLenghChatInputRun(s *discordgo.Session, i *discordgo.InteractionCreate) {
	c.dataLengthRun(s, i)
}
