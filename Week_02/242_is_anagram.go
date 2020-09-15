package Week_02

func isAnagram(s string, t string) bool {
	dict := make(map[int]int, 128)
	for _, c := range s {
		dict[int(c)] = dict[int(c)] + 1
	}

	for _, c := range t {
		dict[int(c)] = dict[int(c)] - 1
	}

	for _, value := range dict {
		if value != 0 {
			return false
		}
	}

	return true
}
