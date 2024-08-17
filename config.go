package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	token string
}

// LoadConfig a config
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("[goMuffin] 봇의 설절파일을 불러올 수가 없어요.")
		log.Fatalln(err)
	}
	config := Config{}
	setConfig(&config)

	return &config
}

func setConfig(config *Config) {
	token := os.Getenv("TOKEN")
	config.token = token
}
