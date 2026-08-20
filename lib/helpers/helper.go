package helpers

import (
	"unicode"

	"github.com/forPelevin/gomoji"
)

func HasUnwantedChar(category, s string) bool {
	switch category {
	case "username":
		return len(s) != 0 && !hasSymbol(s) && !gomoji.ContainsEmoji(s)
	case "email":
		return len(s) != 0 && !hasSymbol(s) && !gomoji.ContainsEmoji(s) // PERF:&& strings.Contains(s, "@gmail.com") inlogical a 3chiri
	default:
		return false
	}
}

func hasSymbol(str string) bool {
	for _, r := range str {
		if unicode.IsSymbol(r) {
			return true
		}
	}
	return false
}
