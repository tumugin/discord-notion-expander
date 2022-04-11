package main

import (
	"discord-notion-expander/app"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := app.GetConfig()
	if err != nil {
		log.Fatal(err)
		return
	}

	// Required Discord permission: 121856
	discord, err := discordgo.New("Bot " + config.DiscordToken)
	if err != nil {
		log.Fatal(err)
		return
	}

	messageEventHandler := app.NewMessageEventHandler(config.BaseNotionUrl, config.NotionToken)
	discord.AddHandler(messageEventHandler.OnMessage)

	if err := discord.Open(); err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	if err := discord.Close(); err != nil {
		log.Fatal(err)
		return
	}
}
