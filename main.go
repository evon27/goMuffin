package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"gitlab.wh64.net/muffin/goMuffin/configs"
	"gitlab.wh64.net/muffin/goMuffin/handler"
)

func main() {
	config := configs.Config

	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("[goMuffin] 봇의 세션을 만들수가 없어요.")
		log.Fatalln(err)
	}

	dg.AddHandler(handler.MessageCreate)

	dg.Open()

	fmt.Println("[goMuffin] 봇이 실행되고 있어요. 버전:", configs.MUFFIN_VERSION)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
