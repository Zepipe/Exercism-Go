package atbash

import ("unicode"
        "strings"
       )

func Atbash(s string) string {
	if len(s) == 0 {
        return ""
    }

   backwardsAlphabet := map[int]string {
       25: "a", 24: "b", 23: "c", 22: "d",
       21: "e", 20: "f", 19: "g", 18: "h",
       17: "i", 16: "j", 15: "k", 14: "l",
       13: "m", 12: "n", 11: "o", 10: "p",
        9: "q",  8: "r",  7: "s",  6: "t",
        5: "u",  4: "v",  3: "w",  2: "x",
        1: "y", 0 : "z",
    }
    
    result, shift := "", 0
	for _, letter := range strings.ToLower(s) {
        switch {
            case unicode.IsLetter(letter):
            	result += backwardsAlphabet[int(letter - 'a')]
            	shift++
            case unicode.IsDigit(letter):
            	result += string(letter)
            	shift++
            default: // punctuation
            	continue
		}
        // every 5 letters should be together + whitespace
        if shift % 5 == 0 {
            result += " "
        }
    }

    if result[len(result) - 1] == ' ' {
        return result[:len(result) - 1] // to skip last whitespace
    }
    return result
}
