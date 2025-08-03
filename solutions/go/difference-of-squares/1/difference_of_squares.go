package diffsquares

func SquareOfSum(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	return res*res
}

func SumOfSquares(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		res += i * i
	}
	return res
}

func Difference(n int) int {
	diff := SquareOfSum(n) - SumOfSquares(n)
	return diff
}
