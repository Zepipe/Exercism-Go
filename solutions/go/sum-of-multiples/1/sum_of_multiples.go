package summultiples

func SumMultiples(limit int, divisors ...int) int {
    if limit == 0 {
        return 0
    }
    
    uniqueValues := make(map[int]bool)
    
    for _, divisor := range divisors {
        // Skip zero or negative divisors (they'd cause infinite loops or no multiples)
        if divisor <= 0 {
            continue
        }
        
        // Start from the divisor itself
        multiple := divisor
        
        // Find all multiples less than the limit
        for multiple < limit {
            uniqueValues[multiple] = true
            multiple += divisor
        }
    }
    
    // Sum all unique multiples
    sum := 0
    for multiple := range uniqueValues {
        sum += multiple
    }
    return sum
}