package diffsquares

func SquareOfSums(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i
	}

	return sum * sum
}

func SumOfSquares(n int) (squares int) {
	for i := 1; i <= n; i++ {
		squares += i * i
	}

	return
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
