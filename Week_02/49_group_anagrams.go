package Week_02

func groupAnagrams(strs []string) [][]string {
	if len(strs) < 2 {
		return [][]string{strs}
	}
	flags := [26]int64{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	dict := make(map[int64][]string, 1)
	aValue := int('a')

	for _, str := range strs {
		bytes := []byte(str)
		flag := int64(1)
		for _, b := range bytes {
			flag *= flags[int(b)-aValue]
		}

		if dict[flag] == nil {
			dict[flag] = make([]string, 0, 1)
		}
		dict[flag] = append(dict[flag], str)
	}

	result := make([][]string, 0, len(dict))
	for _, childStrs := range dict {
		result = append(result, childStrs)
	}

	return result
}
