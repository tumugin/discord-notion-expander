package main

import (
	"discord-notion-expander/utils"
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/kjk/notionapi"
	"golang.org/x/exp/utf8string"
)

type MessageEventHandler struct {
	baseUrl        string
	notionApiToken string
}

func NewMessageEventHandler(baseUrl string, notionApiToken string) *MessageEventHandler {
	return &MessageEventHandler{
		baseUrl:        baseUrl,
		notionApiToken: notionApiToken,
	}
}

func (messageEventHandler *MessageEventHandler) onMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	match, notionPageId := utils.GetNotionPageIdFromMessage(messageEventHandler.baseUrl, message.Message.Content)
	if !match {
		return
	}
	client := &notionapi.Client{}
	client.AuthToken = messageEventHandler.notionApiToken
	page, err := client.DownloadPage(notionPageId)
	if err != nil {
		log.Println(err)
		return
	}
	rootPage := page.Root()
	title := rootPage.Title
	pageText := utils.GetNotionTextFromBlocks(rootPage.Content)
	pageTextWithMaxLength := utf8string.NewString(pageText).Slice(0, 250)
	if _, err := session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		URL:         page.NotionURL(),
		Title:       title,
		Description: pageTextWithMaxLength,
		Provider: &discordgo.MessageEmbedProvider{
			URL:  "https://www.notion.so/",
			Name: "Notion",
		},
	}); err != nil {
		log.Println(err)
	}
}
