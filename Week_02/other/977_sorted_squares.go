package other

import "math"

func sortedSquares(a []int) []int {
	newArr := make([]int, 0, len(a))
	i, j := -1, math.MaxInt64
	for index, value := range a {
		if value < 0 {
			i = index
		} else {
			j = index
			break
		}
	}

	pos := 0
	for i >= 0 && j < len(a) {
		if a[i]+a[j] < 0 {
			newArr[pos] = a[j] * a[j]
			j++
		} else {
			newArr[pos] = a[i] * a[i]
			i--
		}
		pos++
	}

	for i >= 0 {
		newArr[pos] = a[i] * a[i]
		i--
		pos++
	}

	for j < len(a) {
		newArr[pos] = a[j] * a[j]
		j++
		pos++
	}
	return newArr
}
