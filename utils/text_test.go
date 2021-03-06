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
		testText: "γ",
		result:   false,
	},
	{
		testText: "ε½",
		result:   false,
	},
	{
		testText: "π",
		result:   true,
	},
	{
		testText: "β€οΈ",
		result:   true,
	},
	{
		testText: "πΆβπ«οΈ",
		result:   true,
	},
}

func TestIsSingleEmojiText(t *testing.T) {
	for _, testItem := range paramsTestIsSingleEmojiText {
		result := IsSingleEmojiText(testItem.testText)
		assert.Equal(t, testItem.result, result, testItem.testText)
	}
}
