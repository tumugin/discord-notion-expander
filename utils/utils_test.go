package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetNotionUrlRegex(t *testing.T) {
	assert.True(t, getNotionUrlRegex("https://www.notion.so/test/") == "https://www.notion.so/test/(.*-([a-zA-Z0-9]*)|([a-zA-Z0-9]*))")
	assert.True(t, getNotionUrlRegex("https://www.notion.so/test") == "https://www.notion.so/test/(.*-([a-zA-Z0-9]*)|([a-zA-Z0-9]*))")
}

var paramsTestGetNotionPageIdFromMessage = []struct {
	baseUrl string
	content string
	pageIds []string
}{
	{
		baseUrl: "https://www.notion.so/test/",
		content: "https://www.notion.so/test/Test-page-all-c969c9455d7c4dd79c7f860f3ace6429",
		pageIds: []string{"c969c9455d7c4dd79c7f860f3ace6429"},
	},
	{
		baseUrl: "https://www.notion.so/test/",
		content: "https://www.notion.so/test/c969c9455d7c4dd79c7f860f3ace6429",
		pageIds: []string{"c969c9455d7c4dd79c7f860f3ace6429"},
	},
	{
		baseUrl: "https://www.notion.so/test/",
		content: "https://www.notion.so/test/c969c9455d7c4dd79c7f860f3ace6429\nhttps://www.notion.so/test/abcdefg12345abcdefg12345",
		pageIds: []string{"c969c9455d7c4dd79c7f860f3ace6429", "abcdefg12345abcdefg12345"},
	},
	{
		baseUrl: "https://www.notion.so/test/",
		content: "うえしゃまぁああああああ",
		pageIds: []string{},
	},
}

func TestGetNotionPageIdFromMessage(t *testing.T) {
	for _, testItem := range paramsTestGetNotionPageIdFromMessage {
		pageIds := GetNotionPageIdsFromMessage(testItem.baseUrl, testItem.content)
		assert.ElementsMatch(t, testItem.pageIds, pageIds)
	}
}
