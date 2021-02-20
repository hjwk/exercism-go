package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram returns true if a string is an isogram.
func IsIsogram(s string) bool {
	occurences := map[rune]int{}

	for _, r := range strings.ToLower(s) {
		if !unicode.IsLetter(r) {
			continue
		}

		if occurences[r] > 0 {
			return false
		}
		occurences[r]++
	}

	return true
}
