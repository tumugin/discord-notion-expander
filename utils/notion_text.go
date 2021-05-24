package utils

import (
	"github.com/dstotijn/go-notion"
)

func GetNotionTextFromBlocks(blocks []notion.Block) string {
	notionText := ""
	for _, block := range blocks {
		if block.Paragraph != nil {
			notionText += RichTextsToString(block.Paragraph.Text) + "\n"
		}
		if block.Heading1 != nil {
			notionText += "# " + RichTextsToString(block.Heading1.Text) + "\n"
		}
		if block.Heading2 != nil {
			notionText += "## " + RichTextsToString(block.Heading2.Text) + "\n"
		}
		if block.Heading3 != nil {
			notionText += "### " + RichTextsToString(block.Heading3.Text) + "\n"
		}
	}
	return notionText
}

func RichTextsToString(richTexts []notion.RichText) string {
	result := ""
	for _, richText := range richTexts {
		result += richText.PlainText
	}
	return result
}
