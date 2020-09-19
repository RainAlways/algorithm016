package other

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	byteDict := make(map[byte]string, len(pattern))
	strDict := make(map[string]byte, len(pattern))

	for index := range pattern {
		strValue, ok1 := byteDict[pattern[index]]
		_, ok2 := strDict[words[index]]
		if ok1 != ok2 {
			return false
		} else if ok1 == true && strValue != words[index] {
			return false
		} else {
			byteDict[pattern[index]] = words[index]
			strDict[words[index]] = pattern[index]
		}
	}

	return true
}
