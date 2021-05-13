package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var paramsTestIsSingleEmojiText = []struct {
	testText string
	result   bool
}{
	{
		testText: "https://example.com/",
		result:   false,
	},
	{
		testText: "a",
		result:   false,
	},
	{
		testText: "あ",
		result:   false,
	},
	{
		testText: "彅",
		result:   false,
	},
	{
		testText: "😇",
		result:   true,
	},
	{
		testText: "❤️",
		result:   true,
	},
	{
		testText: "😶‍🌫️",
		result:   true,
	},
}

func TestIsSingleEmojiText(t *testing.T) {
	for _, testItem := range paramsTestIsSingleEmojiText {
		result := IsSingleEmojiText(testItem.testText)
		assert.Equal(t, testItem.result, result, testItem.testText)
	}
}
