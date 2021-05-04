package main

import (
	"errors"
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
	if err := godotenv.Load(fmt.Sprintf("./%s.env", os.Getenv("GO_ENV"))); err != nil {
		log.Println(err)
	}
	config := &Config{
		baseNotionUrl: os.Getenv("BASE_NOTION_URL"),
		notionToken:   os.Getenv("NOTION_TOKEN"),
		discordToken:  os.Getenv("DISCORD_TOKEN"),
	}
	if err := checkConfig(config); err != nil {
		return nil, err
	}
	return config, nil
}

func checkConfig(config *Config) error {
	if config.discordToken == "" {
		return errors.New("env value DISCORD_TOKEN is required")
	}
	if config.notionToken == "" {
		return errors.New("env value NOTION_TOKEN is required")
	}
	if config.baseNotionUrl == "" {
		return errors.New("env value BASE_NOTION_URL is required")
	}
	return nil
}
