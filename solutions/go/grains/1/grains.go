package grains

import ("fmt"
        "math"
       )

func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
        return 0, fmt.Errorf("Invalid chessboard square")
    }
    
    return uint64(math.Pow(2, float64(number - 1))), nil
}

func Total() uint64 {
    return ^uint64(0)
}
