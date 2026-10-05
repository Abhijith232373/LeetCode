func scoreOfParentheses(s string) int {
	stack := []int{0}

	for _, ch := range s {
		if ch == '(' {
			stack = append(stack, 0)
		} else {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if last == 0 {
				last = 1
			} else {
				last = 2 * last
			}

			stack[len(stack)-1] += last
		}
	}

	return stack[0]
}