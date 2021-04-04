package luhn

import (
	"strings"
	"unicode"
)

// Valid validates the input number against luhn algorithm
func Valid(nums string) bool {
	var sum int

	// remove all whitespaces
	nums = strings.ReplaceAll(nums, " ", "")

	// single digit input is invalid
	if len(nums) <= 1 {
		return false
	}

	// set the value of the double flg for the first element
	// for odd lenght of string, the first element should not be doubled.
	double := len(nums)%2 == 0

	for _, num := range nums {

		if !unicode.IsDigit(num) {
			return false
		}

		// get the int value from the unicoded value
		digit := int(num - '0')

		// double the number only if it is the second digit
		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		// flip the value so we only double the second digit starting from the right of the input
		double = !double

		// increment the number of the digit being processed
		sum += digit
	}
	return sum%10 == 0
}
