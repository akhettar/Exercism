// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	result := []string{}

	// return an empty array
	if len(rhyme) == 0 {
		return rhyme
	}

	for i := 1; i < len(rhyme); i++ {
		result = append(result, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i-1], rhyme[i]))
	}
	return append(result, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
