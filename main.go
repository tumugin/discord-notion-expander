package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := getConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	// 必要な権限: 121856
	discord, err := discordgo.New("Bot " + config.discordToken)
	if err != nil {
		log.Fatal(err)
		return
	}
	messageEventHandler := NewMessageEventHandler(config.baseNotionUrl, config.notionToken)
	discord.AddHandler(messageEventHandler.onMessage)
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
