package other

func sortArrayByParity(a []int) []int {
	for i, j := 0, 0; i < len(a) && j < len(a); i++ {
		if a[i]%2 != 0 {
			if j < i {
				j = i + 1
			}
			for ; j < len(a); j++ {
				if a[j]%2 == 0 {
					a[i], a[j] = a[j], a[i]
					break
				}
			}
		}
	}

	return a
}
