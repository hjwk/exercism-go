package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if a string is an isogram.
func IsIsogram(s string) bool {
	occurences := map[rune]bool{}

	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) {
			continue
		}

		if occurences[r] {
			return false
		}
		occurences[r] = true
	}

	return true
}
