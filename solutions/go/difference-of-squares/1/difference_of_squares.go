package diffsquares

func SquareOfSum(n int) int {
    var resultSum int
    
	if n == 0 {
        return 0
    }

    for i := 1; i <= n; i++ {
        resultSum = resultSum + i
    }
    return resultSum * resultSum
}

func SumOfSquares(n int) int {
	var resultSum int
    
	if n == 0 {
        return 0
    }

    for i := 1; i <= n; i++ {
        resultSum = resultSum + i * i
    }
    return resultSum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
