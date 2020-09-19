package other

func peakIndexInMountainArray(arr []int) int {
	lo, hi := 0, len(arr)-1
	for lo < hi {
		mi := lo + (hi-lo)/2
		if arr[mi] < arr[mi+1] {
			lo = mi + 1
		} else {
			hi = mi
		}
	}

	return lo
}
