package series

// All returns a list of all contiguous substrings of length n from the string s
func All(n int, s string) []string {
    if n > len(s) {
        return []string{}
    }
    
    result := make([]string, 0, len(s)-n+1)
    for i := 0; i <= len(s)-n; i++ {
        result = append(result, s[i:i+n])
    }
    return result
}

// UnsafeFirst returns the first contiguous substring of length n from the string s
func UnsafeFirst(n int, s string) string {
    if n > len(s) {
        return ""
    }
    return s[:n]
}