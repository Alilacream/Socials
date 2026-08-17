package helpers

import (
	"unicode"

	"github.com/forPelevin/gomoji"
)

func HasUnwantedChar(category, s string) bool {
	if category == "username" {
		return len(s) != 0 && !hasSymbol(s) && !gomoji.ContainsEmoji(s)
	} else if category == "email" {
		return len(s) != 0 && !hasSymbol(s) && !gomoji.ContainsEmoji(s) // PERF:&& strings.Contains(s, "@gmail.com") inlogical a 3chiri
	}
	// as if i passed something else entirely
	return true
}

func hasSymbol(str string) bool {
	for _, r := range str {
		if unicode.IsSymbol(r) {
			return true
		}
	}
	return false
}
