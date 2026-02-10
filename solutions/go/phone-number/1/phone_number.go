package phonenumber

import "fmt"

// Number cleans up a phone number and returns it as a 10-digit string
func Number(phoneNumber string) (string, error) {
	// Remove all non-digit characters
	clean := ""
	for _, ch := range phoneNumber {
		if ch >= '0' && ch <= '9' {
			clean += string(ch)
		}
	}

	// Check for valid length
	if len(clean) == 11 && clean[0] == '1' {
		clean = clean[1:]
	} else if len(clean) != 10 {
		return "", fmt.Errorf("invalid phone number length: %s", phoneNumber)
	}

	// Validate N format (N = 2-9)
	if clean[0] < '2' || clean[0] > '9' {
		return "", fmt.Errorf("area code must start with digit 2-9: %s", phoneNumber)
	}
	if clean[3] < '2' || clean[3] > '9' {
		return "", fmt.Errorf("exchange code must start with digit 2-9: %s", phoneNumber)
	}

	return clean, nil
}

// AreaCode extracts the area code from a phone number
func AreaCode(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return num[:3], nil
}

// Format formats a phone number as (NXX) NXX-XXXX
func Format(phoneNumber string) (string, error) {
	num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", num[:3], num[3:6], num[6:]), nil
}