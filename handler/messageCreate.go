package handler

import (
	"context"
	"log"
	"math/rand"
	"strings"

	"git.wh64.net/muffin/goMuffin/commands"
	"git.wh64.net/muffin/goMuffin/configs"
	"git.wh64.net/muffin/goMuffin/databases"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/v2/bson"
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
			s.ChannelTyping(m.ChannelID)

			var datas []databases.Text
			var filter bson.D

			channel, _ := s.Channel(m.ChannelID)
			if channel.NSFW {
				filter = bson.D{{}}
			} else {
				filter = bson.D{{Key: "persona", Value: "muffin"}}
			}

			cur, err := databases.Texts.Find(context.TODO(), filter)
			if err != nil {
				log.Fatalln(err)
			}

			defer cur.Close(context.Background())
			cur.All(context.TODO(), &datas)

			s.ChannelMessageSendReply(m.ChannelID, datas[rand.Intn(len(datas))].Text, m.Reference())
		}

		commands.Discommand.MessageRun(command, s, m)
	} else {
		return
	}
}
