package task1

func isValid(s string) bool {
	pair := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if right, exists := pair[ch]; exists {
			if len(stack) == 0 || stack[len(stack)-1] != right {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}
