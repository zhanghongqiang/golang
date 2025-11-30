package task1

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		prefix = findPrefix(prefix, strs[i])
	}

	return prefix
}

func findPrefix(str1, str2 string) string {
	min := minLength(len(str1), len(str2))

	index := 0

	for index < min && str1[index] == str2[index] {
		index++
	}

	return str1[:index]
}

func minLength(a, b int) int {
	if a < b {
		return a
	}
	return b
}
