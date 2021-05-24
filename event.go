package main

import (
	"context"
	"discord-notion-expander/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/dstotijn/go-notion"
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
	client := notion.NewClient(messageEventHandler.notionApiToken)
	for _, notionPageId := range notionPageIds {
		postNotionPage(session, message, client, notionPageId)
	}
}

func postNotionPage(session *discordgo.Session, message *discordgo.MessageCreate, client *notion.Client, notionPageId string) {
	page, err := client.FindPageByID(context.Background(), notionPageId)
	if err != nil {
		log.Println(err)
		return
	}

	// TODO: Notion APIに実装されていないのでアイコンは取得できない。実装されたら入れる。
	var title string
	if pageProps, res := page.Properties.(notion.PageProperties); res {
		title = utils.RichTextsToString(pageProps.Title.Title)
	}
	if pageProps, res := page.Properties.(notion.DatabasePageProperties); res {
		if value, kres := pageProps["Name"]; kres {
			title = utils.RichTextsToString(value.Title)
		} else {
			log.Printf("Page title property not found in DatabasePage. Page id = %s.\n", notionPageId)
			return
		}
	}

	contents, err := client.FindBlockChildrenByID(context.Background(), notionPageId, &notion.PaginationQuery{})
	if err != nil {
		log.Println(err)
		return
	}
	pageText := utils.GetNotionTextFromBlocks(contents.Results)
	pageTextUtf8String := utf8string.NewString(pageText)
	pageTextWithMaxLength := pageTextUtf8String.Slice(0, funk.MinInt([]int{250, pageTextUtf8String.RuneCount()}).(int))
	if _, err := session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		URL:         "https://www.notion.so/" + notionPageId,
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
