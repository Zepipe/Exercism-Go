package bob

import (
	"strings"
	"unicode"
)

func Hey(remark string) string {
	// Trim whitespace from the remark
	trimmedRemark := strings.TrimSpace(remark)
	
	if trimmedRemark == "" {
		return "Fine. Be that way!"
	}
	
	// Check if the remark is yelling (contains letters and all are uppercase)
	isYelling := false
	hasLetters := false
	for _, r := range trimmedRemark {
		if unicode.IsLetter(r) {
			hasLetters = true
			if !unicode.IsUpper(r) {
				isYelling = false
				break
			}
			isYelling = true
		}
	}
	isYelling = isYelling && hasLetters
	
	// Check if it's a question (ends with ? after trimming)
	isQuestion := strings.HasSuffix(trimmedRemark, "?")
	
	if isQuestion && isYelling {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion {
		return "Sure."
	} else if isYelling {
		return "Whoa, chill out!"
	} else {
		return "Whatever."
	}
}