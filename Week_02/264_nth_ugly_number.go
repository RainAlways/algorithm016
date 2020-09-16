package Week_02

import "container/heap"

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

func nthUglyNumber2(n int) int {
	h := IntHeap(make([]int, 0, 2*n+1))
	heap.Init(&h)
	heap.Push(&h, 1)

	dict := make(map[int]bool, 3*n)

	nums := make([]int, 0, n)

	steps := [3]int{2, 3, 5}

	for len(nums) != n {
		currentNum := heap.Pop(&h).(int)
		nums = append(nums, currentNum)

		for _, step := range steps {
			if !dict[currentNum*step] {
				dict[currentNum*step] = true
				heap.Push(&h, currentNum*step)
			}
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

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := h.Len() - 1
	result := (*h)[n]
	*h = (*h)[0:n]
	return result
}
