package Week_02

func isAnagram(s string, t string) bool {
	left, right := []byte(s), []byte(t)
	sLen, tLen := len(left), len(right)
	if sLen != tLen {
		return false
	}

	dict := make(map[int]int, 26)
	aVal := int('a')
	for i := 0; i < sLen; i++ {
		dict[int(left[i])-aVal] = dict[int(left[i])-aVal] + 1
		dict[int(right[i])-aVal] = dict[int(right[i])-aVal] - 1
	}

	for _, value := range dict {
		if value != 0 {
			return false
		}
	}

	return true
}
