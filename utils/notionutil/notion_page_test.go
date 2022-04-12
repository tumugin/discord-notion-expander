package notionutil

import (
	"discord-notion-expander/utils"
	"github.com/dstotijn/go-notion"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPageTitleByNotionPageWithEmoji(t *testing.T) {
	actual, _ := GetPageTitleByNotionPage(notion.Page{
		Icon: &notion.Icon{
			Emoji: utils.Ptr("🔥"),
		},
		Properties: notion.PageProperties{
			Title: notion.PageTitle{
				Title: []notion.RichText{{PlainText: "Page Title Test"}},
			},
		},
	})
	assert.Equal(t, "🔥 Page Title Test", actual)
}

func TestGetPageTitleByNotionPage(t *testing.T) {
	actual, _ := GetPageTitleByNotionPage(notion.Page{
		Properties: notion.PageProperties{
			Title: notion.PageTitle{
				Title: []notion.RichText{{PlainText: "Page Title Test"}},
			},
		},
	})
	assert.Equal(t, "Page Title Test", actual)
}

func TestGetPageTitleByNotionDatabasePageWithEmoji(t *testing.T) {
	actual, _ := GetPageTitleByNotionPage(notion.Page{
		Icon: &notion.Icon{
			Emoji: utils.Ptr("🔥"),
		},
		Properties: notion.DatabasePageProperties{
			"test_prop": notion.DatabasePageProperty{
				Title: []notion.RichText{{PlainText: "Page Title Test"}},
			},
		},
	})
	assert.Equal(t, "🔥 Page Title Test", actual)
}
