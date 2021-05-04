package utils

import "github.com/kjk/notionapi"

func GetNotionTextFromBlocks(blocks []*notionapi.Block) string {
	notionText := ""
	for _, block := range blocks {
		for _, content := range block.InlineContent {
			notionText += content.Text + "\n"
		}
	}
	return notionText
}
