package diffsquares

func SquareOfSum(n int) int {
    resultSum := 0
    
    for i := 1; i <= n; i++ {
        resultSum += i
    }
    return resultSum * resultSum
}

func SumOfSquares(n int) int {
	resultSum := 0
    
    for i := 1; i <= n; i++ {
        resultSum += i * i
    }
    return resultSum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
