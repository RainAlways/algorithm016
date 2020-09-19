package other

import "math"

func minWindow(s string, t string) string {
	dict := make(map[uint8]int32, 128)
	for _, v := range t {
		dict[uint8(v)]++
	}

	counter := len(t)
	begin, end, head, d := 0, 0, 0, math.MaxInt64
	for end < len(s) {
		if dict[s[end]] > 0 {
			counter--
		}
		dict[s[end]]--
		end++

		for counter == 0 {
			if end-begin < d {
				d = end - begin
				head = begin
			}

			dict[s[begin]]++
			if dict[s[begin]] > 0 {
				counter++
			}
			begin++
		}
	}

	if d == math.MaxInt64 {
		return ""
	} else {
		return s[head : head+d]
	}
}
