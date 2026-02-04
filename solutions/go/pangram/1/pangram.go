package pangram

import "strings"

func IsPangram(input string) bool {
	var letter string
    
    seen := make(map[string]bool)
    
	if input == "" {
        return false
    }
    
    for _, val := range strings.ToLower(input) {
        letter = string(val)

        if letter == " " || letter == "." || letter == "!" || letter == "," || letter == "_"  || letter == "\"" || letter == "1" || letter == "2" || letter == "3" || letter == "4" || letter == "5" || letter == "6" || letter == "7" || letter == "8" || letter == "9" || letter == "0" {
            continue
        }
        
        if !seen[letter] {
            seen[letter] = true
        } else {
            continue
        }
    }

    return len(seen) == 26
}
