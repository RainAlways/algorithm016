package Week_02

func nthUglyNumber(n int) int {
	step2, step3, step5 := 0, 0, 0
	nums := make([]int, 0, n)
	nums = append(nums, 1)
	for i := 0; i < n-1; i++ {
		num := minInt(minInt(nums[step2]*2, nums[step3]*3), nums[step5]*5)
		nums = append(nums, num)
		if num == nums[step2]*2 {
			step2++
		}
		if num == nums[step3]*3 {
			step3++
		}
		if num == nums[step5]*5 {
			step5++
		}
	}

	return nums[n-1]
}

func minInt(left, right int) int {
	if left < right {
		return left
	}
	return right
}
