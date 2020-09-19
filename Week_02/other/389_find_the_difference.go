package other

func findTheDifference(s string, t string) byte {
	dict := [26]int{}
	for index, c := range s {
		dict[int(c-'a')]++
		dict[int(t[index]-'a')]--
	}

	for i := len(s); i < len(t); i++ {
		dict[int(t[i]-'a')]--
	}

	for index, value := range dict {
		if value < 0 {
			return byte('a' + index)
		}
	}

	return 0
}
