package notionutil

import (
	"discord-notion-expander/utils"
	_ "embed"
	"github.com/dstotijn/go-notion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRichTextsToStringWithEmpty(t *testing.T) {
	actual := RichTextsToString([]notion.RichText{})
	assert.Equal(t, "", actual)
}

func TestRichTextsToString(t *testing.T) {
	actual := RichTextsToString([]notion.RichText{
		{PlainText: "test"},
	})
	assert.Equal(t, "test", actual)
}

//go:embed notion_text_test_input.md
var notionTextExpectedTexts string

func TestGetNotionTextFromBlocks(t *testing.T) {
	actual := GetNotionTextFromBlocks([]notion.Block{
		{
			Heading1: utils.Ptr(
				notion.Heading{
					Text: []notion.RichText{
						{PlainText: "オタクの推しを紹介する"},
					},
				},
			),
		},
		{
			Heading2: utils.Ptr(
				notion.Heading{
					Text: []notion.RichText{
						{PlainText: "村崎ゆうな"},
					},
				},
			),
		},
		{
			Heading3: utils.Ptr(
				notion.Heading{
					Text: []notion.RichText{
						{PlainText: "うなについて紹介する"},
					},
				},
			),
		},
		{
			Paragraph: utils.Ptr(
				notion.RichTextBlock{
					Text: []notion.RichText{
						{PlainText: "うなは群青の世界のアイドルです。かわいいです。"},
					},
				},
			),
		},
		{
			Heading3: utils.Ptr(
				notion.Heading{
					Text: []notion.RichText{
						{PlainText: "うなすき"},
					},
				},
			),
		},
		{
			NumberedListItem: utils.Ptr(
				notion.RichTextBlock{
					Text: []notion.RichText{
						{PlainText: "うなすき大学"},
					},
				},
			),
		},
		{
			NumberedListItem: utils.Ptr(
				notion.RichTextBlock{
					Text: []notion.RichText{
						{PlainText: "ああうな"},
					},
				},
			),
		},
		{
			Heading2: utils.Ptr(
				notion.Heading{
					Text: []notion.RichText{
						{PlainText: "藍井すず"},
					},
				},
			),
		},
		{
			Heading3: utils.Ptr(
				notion.Heading{
					Text: []notion.RichText{
						{PlainText: "あおいすずちゃんかわいいね"},
					},
				},
			),
		},
		{
			Paragraph: utils.Ptr(
				notion.RichTextBlock{
					Text: []notion.RichText{
						{PlainText: "やっぱ藍井すずなんだよなぁ～"},
					},
				},
			),
		},
		{
			Heading3: utils.Ptr(
				notion.Heading{
					Text: []notion.RichText{
						{PlainText: "藍井すずのここが最高"},
					},
				},
			),
		},
		{
			BulletedListItem: utils.Ptr(
				notion.RichTextBlock{
					Text: []notion.RichText{
						{PlainText: "あおいすず"},
					},
				},
			),
		},
		{
			BulletedListItem: utils.Ptr(
				notion.RichTextBlock{
					Text: []notion.RichText{
						{PlainText: "あ～すず"},
					},
				},
			),
		},
		{
			Heading3: utils.Ptr(
				notion.Heading{
					Text: []notion.RichText{
						{PlainText: "藍井すずにふさわしい男になるチェックリスト"},
					},
				},
			),
		},
		{
			ToDo: utils.Ptr(
				notion.ToDo{
					RichTextBlock: notion.RichTextBlock{
						Text: []notion.RichText{
							{PlainText: "あおいすず"},
						},
					},
				},
			),
		},
		{
			Toggle: utils.Ptr(
				notion.RichTextBlock{
					Text: []notion.RichText{
						{PlainText: "あ～すず"},
					},
				},
			),
		},
		{
			ChildPage: utils.Ptr(
				notion.ChildPage{
					Title: "藍井すずふさわしくなるページ",
				},
			),
		},
	})
	assert.Equal(t, notionTextExpectedTexts, actual)
}
