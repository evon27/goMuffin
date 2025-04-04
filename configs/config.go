package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)


type botConfig struct {
	Token   string
	Prefix  string
	OwnerId string
}

type trainConfig struct {
	UserID string
}

// MuffinConfig for Muffin bot
type MuffinConfig struct {
	Bot         botConfig
	Train       trainConfig
	DatabaseURL string
}

func loadConfig() *MuffinConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[goMuffin] 봇의 설절파일을 불러올 수가 없어요.")
		log.Fatalln(err)
	}
	config := MuffinConfig{Bot: botConfig{}, Train: trainConfig{}}
	setConfig(&config)

	return &config
}

func setConfig(config *MuffinConfig) {
	config.Bot.Prefix = os.Getenv("BOT_PREFIX")
	config.Bot.Token = os.Getenv("BOT_TOKEN")
	config.Bot.OwnerId = os.Getenv("BOT_OWNER_ID")

	config.Train.UserID = os.Getenv("TRAIN_USER_ID")

	config.DatabaseURL = os.Getenv("DATABASE_URL")
}

var Config *MuffinConfig = loadConfig()
