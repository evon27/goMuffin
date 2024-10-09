package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var MUFFIN_VERSION = "0.0.0-gopher.e241009a"

// MuffinConfig for Muffin bot
type MuffinConfig struct {
	Token  string
	Prefix string
}

func loadConfig() *MuffinConfig {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[goMuffin] 봇의 설절파일을 불러올 수가 없어요.")
		log.Fatalln(err)
	}
	config := MuffinConfig{}
	setConfig(&config)

	return &config
}

func setConfig(config *MuffinConfig) {
	token := os.Getenv("TOKEN")
	config.Prefix = os.Getenv("PREFIX")
	config.Token = token
}

var Config *MuffinConfig = loadConfig()
