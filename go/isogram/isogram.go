package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram determines if a word or phrase is an isogram.
func IsIsogram(text string) bool {

	// a map holding visited char
	visited := make(map[rune]bool)
	for _, l := range strings.ToLower(text) {

		// check if the current letter has been seen before in the given text
		if visited[l] {
			return false
		}

		// ignore white spaces and hyphens
		if unicode.IsLetter(l) {
			visited[l] = true
		}
	}
	return true
}
