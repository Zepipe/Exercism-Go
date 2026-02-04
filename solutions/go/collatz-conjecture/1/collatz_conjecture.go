package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    var numOfSteps int // To count computation steps
    
    if n <= 0 {
        return numOfSteps, errors.New("Invalid input")
    }
    
    for n > 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n = n * 3 + 1
        }
        numOfSteps++
    }
    return numOfSteps, nil
}
