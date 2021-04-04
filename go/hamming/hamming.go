package hamming

import (
	"fmt"
)


// Distance calculate the hamming distance
// between to DNA strands
func Distance(a, b string) (int, error) {

	// convert the strings to an array of runes
	seqA := []rune(a)
	seqB := []rune(b)

	// The DNA sequence length must be equal
	if len(seqA) != len(seqB) {
		return 0, fmt.Errorf("the length of both sequence must be equal: lenght[a,b]:[%d,%d]", len(a), len(b))
	}

	counter := 0
	for i, val := range seqA {
		if val != seqB[i] {
			counter ++
		}
	}
	return counter, nil
}
