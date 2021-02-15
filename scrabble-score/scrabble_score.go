// Package scrabble provides a scoring function for English Scrabble.
package scrabble

import "strings"

// The size is 128 because we use the ASCII value of a letter as the index and 'z' is 122 in ASCII.
var points = [128]int{
	' ': 0,
	'a': 1,
	'e': 1,
	'i': 1,
	'o': 1,
	'u': 1,
	'l': 1,
	'n': 1,
	'r': 1,
	's': 1,
	't': 1,
	'd': 2,
	'g': 2,
	'b': 3,
	'c': 3,
	'm': 3,
	'p': 3,
	'f': 4,
	'h': 4,
	'v': 4,
	'w': 4,
	'y': 4,
	'k': 5,
	'j': 8,
	'x': 8,
	'q': 10,
	'z': 10,
}

// Score returns the number of points a word is worth, according to English
// Scrabble.
func Score(word string) int {
	var score int
	for _, r := range strings.ToLower(word) {
		score += points[r]
	}
	return score
}
