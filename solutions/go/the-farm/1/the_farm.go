package thefarm

import "fmt"

func DivideFood(fc FodderCalculator, cows int) (float64, error) {
    totalAmount, err := fc.FodderAmount(cows)
    if err != nil {
        return 0, err
    }
    fatFactor, err := fc.FatteningFactor()
    if err != nil {
        return 0, err
    }
    return totalAmount * fatFactor / float64(cows), nil
}

func ValidateInputAndDivideFood(fc FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
        return DivideFood(fc, cows)
    } else {
        return 0, fmt.Errorf("invalid number of cows")
    }
}

// InvalidCowsError is a custom error type for invalid cow numbers
type InvalidCowsError struct {
    NumberOfCows int
    Message      string
}

// Error implements the error interface
func (e *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", e.NumberOfCows, e.Message)
}

// ValidateNumberOfCows checks if the number of cows is valid
// Returns nil if valid, otherwise returns an InvalidCowsError with appropriate message
func ValidateNumberOfCows(cows int) error {
    if cows < 0 {
        return &InvalidCowsError{
            NumberOfCows: cows,
            Message:      "there are no negative cows",
        }
    }
    
    if cows == 0 {
        return &InvalidCowsError{
            NumberOfCows: cows,
            Message:      "no cows don't need food",
        }
    }
    
    return nil
}
