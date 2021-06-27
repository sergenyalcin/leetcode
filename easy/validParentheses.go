package easy

func ValidParentheses(s string) bool {
	stack := ""

	for _, c := range s {
		char := string(c)

		if char == "(" || char == "{" || char == "[" {
			stack += char
			continue
		} else {
			if len(stack) < 1 {
				return false
			}

			if (char == ")" && string(stack[len(stack)-1]) == "(") ||
				(char == "}" && string(stack[len(stack)-1]) == "{") ||
				(char == "]" && string(stack[len(stack)-1]) == "[") {
				stack = stack[0 : len(stack)-1]
			} else {
				return false
			}
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}
