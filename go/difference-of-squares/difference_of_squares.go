package diffsquares

// SquareOfSum return he square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	var sum int
	for i := 0; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares is the sum of the squares of the first ten natural numbers is
func SumOfSquares(n int) int {
	// the base
	if n == 1 {
		return 1
	}
	return n*n + SumOfSquares(n-1)
}

// Difference computes the difference between the square of the sum of the first
// n natural numbers and the sum of the squares of the first n natural numbers
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
