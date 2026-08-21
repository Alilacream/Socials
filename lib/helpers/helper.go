package helpers

import (
	"unicode"

	"github.com/forPelevin/gomoji"
)

// checks if the string is clean or not, specifying the category aswell
func HasCleanChars(category, s string) bool {
	switch category {
	case "username":
		return len(s) >= 8 && !hasSymbol(s) && !gomoji.ContainsEmoji(s)
	case "email":
		return len(s) >= 8 && !hasSymbol(s) && !gomoji.ContainsEmoji(s) // PERF:&& strings.Contains(s, "@gmail.com") inlogical a 3chiri
	case "password":
		return len(s) >= 8 && !hasSymbol(s) && !gomoji.ContainsEmoji(s)
	case "url":
		return s == "" && !hasSymbol(s) && !gomoji.ContainsEmoji(s)
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
