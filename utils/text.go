package utils

import (
	"github.com/rivo/uniseg"
)

func IsSingleEmojiText(text string) bool {
	// 4byte以上の文字なら通しているが、だいぶ雑な判定なので直した方がよさそう
	return uniseg.GraphemeClusterCount(text) == 1 && len(text) > 3
}
