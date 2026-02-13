package prime

func Factors(n int64) []int64 {
	if n <= 1 {
        return []int64{}
    }
    
	var curDivisor int64 = 2
    factors := make([]int64, 0)
    
    for n != 1 {
        if n % curDivisor == 0 {
            n /= curDivisor
            factors = append(factors, curDivisor)
        } else {
            curDivisor++
        }
    }
    return factors
}
