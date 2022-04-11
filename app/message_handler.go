package app

import (
	"context"
	"discord-notion-expander/utils"
	"discord-notion-expander/utils/notionutil"
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

func (messageEventHandler *MessageEventHandler) OnMessage(session *discordgo.Session, message *discordgo.MessageCreate) {
	notionPageIds := utils.GetNotionPageIdsFromMessage(messageEventHandler.baseUrl, message.Message.Content)

	if len(notionPageIds) == 0 {
		// 不要な認証などをかけない為に必要ないときは処理を止める
		return
	}

	client := notion.NewClient(messageEventHandler.notionApiToken)

	for _, notionPageId := range notionPageIds {
		if err := postNotionPage(session, message, client, notionPageId); err != nil {
			log.Println(err)
		}
	}
}

func postNotionPage(session *discordgo.Session, message *discordgo.MessageCreate, client *notion.Client, notionPageId string) error {
	page, err := client.FindPageByID(context.Background(), notionPageId)
	if err != nil {
		return err
	}

	title, err := notionutil.GetPageTitleByNotionPage(page)
	if err != nil {
		return err
	}

	contents, err := client.FindBlockChildrenByID(context.Background(), notionPageId, &notion.PaginationQuery{})
	if err != nil {
		return err
	}

	pageText := notionutil.GetNotionTextFromBlocks(contents.Results)
	pageTextUtf8String := utf8string.NewString(pageText)
	pageTextWithMaxLength := pageTextUtf8String.Slice(
		0,
		funk.MinInt([]int{250, pageTextUtf8String.RuneCount()}),
	)

	if _, err := session.ChannelMessageSendEmbed(message.ChannelID, &discordgo.MessageEmbed{
		URL:         "https://www.notion.so/" + notionPageId,
		Title:       title,
		Description: pageTextWithMaxLength,
		Provider: &discordgo.MessageEmbedProvider{
			URL:  "https://www.notion.so/",
			Name: "Notion",
		},
	}); err != nil {
		return err
	}

	return nil
}
