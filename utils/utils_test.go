package utils

import (
	"github.com/alecthomas/assert"
	"testing"
)

func TestGetNotionUrlRegex(t *testing.T) {
	assert.True(t, getNotionUrlRegex("https://www.notion.so/test/") == "https://www.notion.so/test/(.*-([a-zA-Z0-9]*)|([a-zA-Z0-9]*))")
	assert.True(t, getNotionUrlRegex("https://www.notion.so/test") == "https://www.notion.so/test/(.*-([a-zA-Z0-9]*)|([a-zA-Z0-9]*))")
}

var paramsTestGetNotionPageIdFromMessage = []struct {
	baseUrl string
	content string
	match   bool
	pageId  string
}{
	{
		baseUrl: "https://www.notion.so/test/",
		content: "https://www.notion.so/test/Test-page-all-c969c9455d7c4dd79c7f860f3ace6429",
		match:   true,
		pageId:  "c969c9455d7c4dd79c7f860f3ace6429",
	},
	{
		baseUrl: "https://www.notion.so/test/",
		content: "https://www.notion.so/test/c969c9455d7c4dd79c7f860f3ace6429",
		match:   true,
		pageId:  "c969c9455d7c4dd79c7f860f3ace6429",
	},
	{
		baseUrl: "https://www.notion.so/test/",
		content: "うえしゃまぁああああああ",
		match:   false,
		pageId:  "",
	},
}

func TestGetNotionPageIdFromMessage(t *testing.T) {
	for _, testItem := range paramsTestGetNotionPageIdFromMessage {
		match, pageId := GetNotionPageIdFromMessage(testItem.baseUrl, testItem.content)
		assert.Equal(t, testItem.match, match)
		if testItem.match {
			assert.Equal(t, testItem.pageId, pageId)
		}
	}
}
