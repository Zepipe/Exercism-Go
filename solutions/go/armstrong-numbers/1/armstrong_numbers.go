package armstrong

import ("strconv"
        "unicode"
        "math")

func IsNumber(n int) bool {
	var armstrongNum, countDigits float64 
	strNum := strconv.Itoa(n)
    
    for _, letter := range strNum {
        if !unicode.IsDigit(letter) {
            return false
        }
        countDigits++
    }

    for _, val := range strNum {
        armstrongNum += math.Pow(float64(val - '0'), countDigits)
    }

    return int(armstrongNum) == n
}
