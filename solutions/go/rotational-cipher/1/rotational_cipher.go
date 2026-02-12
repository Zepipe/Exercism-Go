package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
    if plain == "" {
        return ""
    }
    
    // Normalize shift key to 0-25 range using modulo
    shiftKey = shiftKey % 26
    
    // No shift needed
    if shiftKey == 0 {
        return plain
    }
    
    result := make([]rune, 0, len(plain))
    
    for _, letter := range plain {
        switch {
        case unicode.IsUpper(letter):
            // Shift uppercase letter with wrap around
            shifted := 'A' + (letter-'A'+rune(shiftKey))%26
            result = append(result, shifted)
        case unicode.IsLower(letter):
            // Shift lowercase letter with wrap around
            shifted := 'a' + (letter-'a'+rune(shiftKey))%26
            result = append(result, shifted)
        default:
            // Non-letters remain unchanged
            result = append(result, letter)
        }
    }
    
    return string(result)
}