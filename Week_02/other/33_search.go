package other

func search(nums []int, target int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	if n == 1 {
		if target == nums[0] {
			return 0
		} else {
			return -1
		}
	}

	lo, hi := 0, n-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		if target == nums[mid] {
			return mid
		}

		if nums[lo] < nums[mid] {
			if nums[lo] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if nums[mid+1] <= target && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}

	if nums[lo] == target {
		return lo
	}

	return -1
}
