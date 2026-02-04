package reverse

func Reverse(input string) string {
	var reversedStr string

	if input == "" {
		return ""
	}

	for _, val := range input {
		reversedStr = string(val) + reversedStr
	}
	return reversedStr
}
