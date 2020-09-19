package other

func firstUniqChar(s string) int {
	dict := [26]int{}
	for _, c := range s {
		dict[c-'a']++
	}

	for index, c := range s {
		if dict[c-'a'] == 1 {
			return index
		}
	}

	return -1
}
