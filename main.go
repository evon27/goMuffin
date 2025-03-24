package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"git.wh64.net/muffin/goMuffin/configs"
	"git.wh64.net/muffin/goMuffin/databases"
	"git.wh64.net/muffin/goMuffin/handler"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config := configs.Config

	dg, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		fmt.Println("[goMuffin] 봇의 세션을 만들수가 없어요.")
		log.Fatalln(err)
	}

	dg.AddHandler(handler.MessageCreate)

	dg.Open()

	defer func() {
		if err := databases.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	fmt.Println("[goMuffin] 봇이 실행되고 있어요. 버전:", configs.MUFFIN_VERSION)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
