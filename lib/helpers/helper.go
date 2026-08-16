package helpers

import (
	"strings"
	"unicode"

	"github.com/forPelevin/gomoji"
)

func HasUnwantedChar(s string) bool {
	return len(s) == 0 || strings.Contains(s, "@gmail.com") || hasSymbol(s) || gomoji.ContainsEmoji(s)
}

func hasSymbol(str string) bool {
	for _, r := range str {
		if unicode.IsSymbol(r) {
			return true
		}
	}
	return false
}
