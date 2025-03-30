package handler

import (
	"context"
	"log"
	"math/rand"
	"strings"
	"time"

	"git.wh64.net/muffin/goMuffin/commands"
	"git.wh64.net/muffin/goMuffin/configs"
	"git.wh64.net/muffin/goMuffin/databases"
	"git.wh64.net/muffin/goMuffin/utils"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func argParser(content string) (args []string) {
	for _, arg := range utils.FlexibleStringParser.FindAllStringSubmatch(content, -1) {
		if arg[1] != "" {
			args = append(args, arg[1])
		} else {
			args = append(args, arg[0])
		}
	}
	return
}

// MessageCreate is handlers of messageCreate event
func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	config := configs.Config
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	if strings.HasPrefix(m.Content, config.Bot.Prefix) {
		content := strings.TrimPrefix(m.Content, config.Bot.Prefix)
		args := argParser(content)
		command := commands.Discommand.Aliases[args[0]]

		if m.Author.ID == config.Train.UserID {
			if _, err := databases.Texts.InsertOne(context.TODO(), databases.InsertText{
				Text:      content,
				Persona:   "muffin",
				CreatedAt: time.Now(),
			}); err != nil {
				log.Fatalln(err)
			}
		}

		if command == "" {
			s.ChannelTyping(m.ChannelID)

			var datas []databases.Text
			var learnDatas []databases.Learn
			var filter bson.D

			ch := make(chan int)
			x := rand.Intn(5)

			channel, _ := s.Channel(m.ChannelID)
			if channel.NSFW {
				filter = bson.D{{}}

				if _, err := databases.Texts.InsertOne(context.TODO(), databases.InsertText{
					Text:      content,
					Persona:   "user:" + m.Author.Username,
					CreatedAt: time.Now(),
				}); err != nil {
					log.Fatalln(err)
				}

			} else {
				filter = bson.D{{Key: "persona", Value: "muffin"}}
			}

			go func() {
				cur, err := databases.Texts.Find(context.TODO(), filter)
				if err != nil {
					log.Fatalln(err)
				}

				defer cur.Close(context.TODO())

				cur.All(context.TODO(), &datas)
				ch <- 1
			}()
			go func() {
				cur, err := databases.Learns.Find(context.TODO(), bson.D{{Key: "command", Value: content}})
				if err != nil {
					if err == mongo.ErrNilDocument {
						learnDatas = []databases.Learn{}
					}
					log.Fatalln(err)
				}

				defer cur.Close(context.TODO())

				cur.All(context.TODO(), &learnDatas)
				ch <- 1
			}()

			for range 2 {
				<-ch
			}

			if x > 2 && len(learnDatas) != 0 {
				data := learnDatas[rand.Intn(len(learnDatas))]
				user, _ := s.User(data.UserId)
				s.ChannelMessageSendReply(m.ChannelID, data.Result+"\n"+utils.InlineCode(user.Username+"님이 알려주셨어요."), m.Reference())
				return
			}

			s.ChannelMessageSendReply(m.ChannelID, datas[rand.Intn(len(datas))].Text, m.Reference())
			return
		}

		commands.Discommand.MessageRun(command, s, m, args[1:])
		return
	} else {
		return
	}
}
