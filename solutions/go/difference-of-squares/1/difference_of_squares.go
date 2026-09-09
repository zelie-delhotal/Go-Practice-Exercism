// Package differenceofsquares contains one function to calculate the difference
// between the sum of squares and the square of the sum of the first n integers.
package differenceofsquares

func SquareOfSum(n int) int {
    var sum int
    for i:=1; i<=n; i++ {
        sum += i
    }
    return sum * sum
}

func SumOfSquares(n int) int {
    var sum int
    for i:=1; i<=n; i++ {
        sum += i * i
    }
    return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
