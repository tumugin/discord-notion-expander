package app

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	BaseNotionUrl string
	NotionToken   string
	DiscordToken  string
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV"))); err != nil {
		log.Println(err)
	}
	config := &Config{
		BaseNotionUrl: os.Getenv("BASE_NOTION_URL"),
		NotionToken:   os.Getenv("NOTION_TOKEN"),
		DiscordToken:  os.Getenv("DISCORD_TOKEN"),
	}
	if err := checkConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

func checkConfig(config *Config) error {
	if config.DiscordToken == "" {
		return errors.New("env value DISCORD_TOKEN is required")
	}
	if config.NotionToken == "" {
		return errors.New("env value NOTION_TOKEN is required")
	}
	if config.BaseNotionUrl == "" {
		return errors.New("env value BASE_NOTION_URL is required")
	}
	return nil
}
