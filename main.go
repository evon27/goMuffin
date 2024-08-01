package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Muffin-laboratory/goMuffin/handler"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	token := os.Getenv("TOKEN")
	if err != nil {
		fmt.Println("[goMuffin] 봇의 설절파일을 불러올 수가 없어요.")
		log.Fatalln(err)
	}

	dg, err := discordgo.New("Bot " + token)
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

	// Cleanly close down the Discord session.
	dg.Close()
}
