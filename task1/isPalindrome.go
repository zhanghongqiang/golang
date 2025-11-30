package task1

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	originNumber := x

	reverseNumber := 0

	for x > 0 {
		reverseNumber = reverseNumber*10 + x%10
		x = int(x / 10)
	}

	return originNumber == reverseNumber
}
