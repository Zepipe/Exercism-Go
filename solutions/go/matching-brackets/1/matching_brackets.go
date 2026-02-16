package brackets

func Bracket(input string) bool {
    stack := []rune{}
    
    for _, char := range input {
        switch char {
        case '{', '[', '(':
            // Push opening brackets onto the stack
            stack = append(stack, char)
            
        case '}':
            // Check if the top of the stack matches
            if len(stack) == 0 || stack[len(stack)-1] != '{' {
                return false
            }
            stack = stack[:len(stack)-1] // Pop from stack
            
        case ']':
            if len(stack) == 0 || stack[len(stack)-1] != '[' {
                return false
            }
            stack = stack[:len(stack)-1]
            
        case ')':
            if len(stack) == 0 || stack[len(stack)-1] != '(' {
                return false
            }
            stack = stack[:len(stack)-1]
        }
        // Ignore all other characters (including spaces)
    }
    
    // If stack is empty, all brackets were properly closed
    return len(stack) == 0
}