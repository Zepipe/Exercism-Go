package hamming

import "errors"

// Distance calculates the hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("length is not equal")
	}
	
	count := 0
	
	for i, s := range []byte(a) {
		if s != b[i] {
			count++
		}
	}
	return count, nil
}
