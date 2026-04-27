
func isValid(s string) bool {
    // Map closing -> opening for easy lookup
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    
    // Stack of runes
    stack := make([]rune, 0)

    for _, char := range s {
        // If char is a key in pairs, it's a closing bracket
        if open, isClosing := pairs[char]; isClosing {
            // Check if stack is empty or top doesn't match
            if len(stack) == 0 || stack[len(stack)-1] != open {
                return false
            }
            // Pop from stack
            stack = stack[:len(stack)-1]
        } else {
            // It's an opening bracket, push to stack
            stack = append(stack, char)
        }
    }

    return len(stack) == 0
}