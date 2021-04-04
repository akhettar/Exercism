package raindrops

import (
	"strconv"
)

//Convert an integer into a string according to the following rules:
// - has 3 as a factor, add 'Pling' to the result.
// - has 5 as a factor, add 'Plang' to the result.
// - has 7 as a factor, add 'Plong' to the result.
func Convert(num int) string {
	var result string

	if num%3 == 0 {
		result += "Pling"
	}
	if num%5 == 0 {
		result += "Plang"
	}
	if num%7 == 0 {
		result += "Plong"
	}

	// No factor for the given number, return the given number
	if result == "" {
		return strconv.Itoa(num)
	}
	return result
}
