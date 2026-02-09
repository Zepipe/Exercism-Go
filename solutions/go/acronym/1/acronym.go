package acronym

import "strings"

func Abbreviate(s string) string {
	result := ""

    for idx, val := range s {
        if idx == 0 {
            result += string(val)
        } else if val == ' ' || val == '-' || val == '_' || val == ',' {
            result += strings.ToUpper(string(s[idx + 1])) // take next element 
        }
    }

    return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(result, " ", ""), "_", ""), "-", "")
}
