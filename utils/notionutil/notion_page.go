package notionutil

import (
	"errors"
	"fmt"
	"github.com/dstotijn/go-notion"
)

func GetPageTitleByNotionPage(page notion.Page) (string, error) {
	// Emoji text
	var prefixEmojiText = ""
	if page.Icon != nil && page.Icon.Emoji != nil {
		prefixEmojiText = *page.Icon.Emoji + " "
	}

	// Normal page
	if pageProps, res := page.Properties.(notion.PageProperties); res {
		return prefixEmojiText + RichTextsToString(pageProps.Title.Title), nil
	}

	// Database page
	if pageProps, res := page.Properties.(notion.DatabasePageProperties); res {
		for _, propertyValue := range pageProps {
			if propertyValue.Title != nil {
				return prefixEmojiText + RichTextsToString(propertyValue.Title), nil
			}
		}
		return "", errors.New(fmt.Sprintf("Page title property not found in DatabasePage. Page id = %s.\n", page.ID))
	}

	return "", errors.New(fmt.Sprintf("Page title property not found in unknown page propperties. Page id = %s.\n", page.ID))
}
