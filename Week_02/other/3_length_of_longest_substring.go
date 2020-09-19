package other

func lengthOfLongestSubstring(s string) int {
	dict := make(map[int32]int, 10)
	max := 0
	j := 0
	for i, v := range s {
		p, ok := dict[v]
		if ok {
			j = maxInt(j, p+1)
		}
		dict[v] = i
		max = maxInt(max, i-j+1)
	}

	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
