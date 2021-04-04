package scrabble

import "strings"

// Score computes the Scrable score for the given word
// The function uses the below values for given letter
// to workout the total score:
//	Letter                           Value
//	A, E, I, O, U, L, N, R, S, T       1
//	D, G                               2
//	B, C, M, P                         3
//	F, H, V, W, Y                      4
//	K                                  5
//	J, X                               8
//	Q, Z                               10
func Score(word string) int {
	var total int
	for _, w := range strings.ToUpper(word) {

		switch w {
		case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
			total++
		case 'D', 'G':
			total += 2
		case 'B', 'C', 'M', 'P':
			total += 3
		case 'F', 'H', 'V', 'W', 'Y':
			total += 4
		case 'K':
			total += 5
		case 'J', 'X':
			total += 8
		case 'Q', 'Z':
			total += 10
		}
	}
	return total
}
