package darts

import "math"

func Score(x, y float64) int {
	doubledX, doubledY := math.Pow(x, 2.0), math.Pow(y, 2.0)
    distanceToRadius := math.Pow(doubledX + doubledY, 0.5)

    switch {
        case distanceToRadius <= 1: // inner circle
            return 10
        case distanceToRadius <= 5: // middle circle
        	return 5
        case distanceToRadius <= 10: // outer circle
        	return 1
        default:
        	return 0
    }
}
