package notionutil

import (
	"errors"
	"fmt"
	"github.com/dstotijn/go-notion"
)

func GetPageTitleByNotionPage(page notion.Page) (string, error) {
	// TODO: Notion APIに実装されていないのでアイコンは取得できない。実装されたら入れる。

	// Normal page
	if pageProps, res := page.Properties.(notion.PageProperties); res {
		return RichTextsToString(pageProps.Title.Title), nil
	}

	// Database page
	if pageProps, res := page.Properties.(notion.DatabasePageProperties); res {
		for _, propertyValue := range pageProps {
			if propertyValue.Title != nil {
				return RichTextsToString(propertyValue.Title), nil
			}
		}
		return "", errors.New(fmt.Sprintf("Page title property not found in DatabasePage. Page id = %s.\n", page.ID))
	}

	return "", errors.New(fmt.Sprintf("Page title property not found in unknown page propperties. Page id = %s.\n", page.ID))
}
