package main

import (
	"fmt"
	"github.com/joho/godotenv"
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
		return nil, err
	}
	config := &Config{
		baseNotionUrl: os.Getenv("BASE_NOTION_URL"),
		notionToken:   os.Getenv("NOTION_TOKEN"),
		discordToken:  os.Getenv("DISCORD_TOKEN"),
	}
	return config, err
}
