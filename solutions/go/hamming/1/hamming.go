package hamming

import "unicode/utf8"

type param_error struct{}
 
func (error_object param_error) Error() string{ 
    return "Invalid parameter"
}

func Distance(a, b string) (int, error) {
    var hammingDistance int
    
	if a == "" && b == "" {
        return 0, nil
    } else if a == "" || b == "" {
        return 0, &param_error{}
    } else if utf8.RuneCountInString(a) != utf8.RuneCountInString(b) {
    	return 0, &param_error{}
    } else if a == b {
        return 0, nil
    }

    for idx, val := range a {
        for idx2, val2 := range b {
            if idx == idx2 { // to find equal indexes
                if string(val) != string(val2) {
                    hammingDistance += 1
                } else {
                    break
                }
            }
        }
    }

    return hammingDistance, nil
    
}
