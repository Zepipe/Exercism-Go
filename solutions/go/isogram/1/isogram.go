package isogram

import "strings"

func IsIsogram(word string) bool {
	var letter string
    seen := make(map[string]bool)
    if word == "" {
        return true
    }

    for _, val := range strings.ToLower(word) {
        letter = string(val) // each letter in the word

        if letter == " " || letter == "-" {
            continue
        }
        
        if !seen[letter] {
            seen[letter] = true
        } else {
            return false // found second letter occurance
        }
    }

	return true
}
