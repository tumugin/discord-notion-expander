package main

import (
	"discord-notion-expander/utils"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/kjk/notionapi"
	"github.com/thoas/go-funk"
	"golang.org/x/exp/utf8string"
	"log"
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
	notionPageIds := utils.GetNotionPageIdsFromMessage(messageEventHandler.baseUrl, message.Message.Content)
	if len(notionPageIds) == 0 {
		// 不要な認証などをかけない為に必要ないときは処理を止める
		return
	}
	client := &notionapi.Client{}
	client.AuthToken = messageEventHandler.notionApiToken
	for _, notionPageId := range notionPageIds {
		postNotionPage(session, message, client, notionPageId)
	}
}

func postNotionPage(session *discordgo.Session, message *discordgo.MessageCreate, client *notionapi.Client, notionPageId string) {
	page, err := client.DownloadPage(notionPageId)
	if err != nil {
		log.Println(err)
		return
	}
	rootPage := page.Root()
	formatPage := rootPage.FormatPage()
	var title string
	if formatPage != nil && utils.IsSingleEmojiText(formatPage.PageIcon) {
		title = fmt.Sprintf("%s %s", formatPage.PageIcon, rootPage.Title)
	} else {
		title = rootPage.Title
	}
	pageText := utils.GetNotionTextFromBlocks(rootPage.Content)
	pageTextUtf8String := utf8string.NewString(pageText)
	pageTextWithMaxLength := pageTextUtf8String.Slice(0, funk.MinInt([]int{250, pageTextUtf8String.RuneCount()}).(int))
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
