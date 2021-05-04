package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	baseNotionUrl string
	notionToken   string
	discordToken  string
}

func getConfig() (*Config, error) {
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Println(err)
	}
	config := &Config{
		baseNotionUrl: os.Getenv("BASE_NOTION_URL"),
		notionToken:   os.Getenv("NOTION_TOKEN"),
		discordToken:  os.Getenv("DISCORD_TOKEN"),
	}
	checkConfig(config)
	return config, err
}

func checkConfig(config *Config) {
	if config.discordToken == "" {
		log.Fatal("Env value DISCORD_TOKEN is required.")
	}
	if config.notionToken == "" {
		log.Fatal("Env value NOTION_TOKEN is required.")
	}
	if config.baseNotionUrl == "" {
		log.Fatal("Env value BASE_NOTION_URL is required.")
	}
}
