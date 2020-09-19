package other

func sortArrayByParityII(a []int) []int {
	cur := 0
	for cur < len(a) {
		if cur%2 != a[cur]%2 {
			for j := cur + 1; j < len(a); j++ {
				if j%2 != a[j]%2 && a[j]%2 == cur%2 {
					a[cur], a[j] = a[j], a[cur]
					break
				}
			}
		}
		cur++
	}

	return a
}

func sortArrayByParityII2(a []int) []int {
	for i, j := 0, 1; i < len(a); i += 2 {
		if a[i]%2 != 0 {
			for j < len(a) {
				if a[j]%2 != 1 {
					a[i], a[j] = a[j], a[i]
					break
				}
				j += 2
			}
		}
	}

	return a
}
