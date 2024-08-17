package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Muffin-laboratory/goMuffin/handler"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config := LoadConfig()

	dg, err := discordgo.New("Bot " + config.token)
	if err != nil {
		fmt.Println("[goMuffin] 봇의 세션을 만들수가 없어요.")
		log.Fatalln(err)
	}

	dg.AddHandler(handler.MessageCreate)

	dg.Open()

	fmt.Println("[goMuffin] 봇이 실행되고 있어요.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
