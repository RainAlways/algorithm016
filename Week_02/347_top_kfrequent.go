package Week_02

import "container/heap"

func topKFrequent(nums []int, k int) []int {
	dict := make(map[int]int, len(nums)/2)

	for _, value := range nums {
		dict[value] = dict[value] + 1
	}

	h := scoreHeap(make([][2]int, 0, k+1))
	heap.Init(&h)
	for value, count := range dict {
		heap.Push(&h, [2]int{value, count})
		if h.Len() > k {
			heap.Pop(&h)
		}
	}

	result := make([]int, k, k)
	for i := 0; i < k; i++ {
		value := heap.Pop(&h).([2]int)
		result[k-i-1] = value[0]
	}

	return result
}

type scoreHeap [][2]int

func (s scoreHeap) Len() int {
	return len(s)
}

func (s scoreHeap) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s scoreHeap) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s *scoreHeap) Push(x interface{}) {
	*s = append(*s, x.([2]int))
}

func (s *scoreHeap) Pop() interface{} {
	n := s.Len() - 1
	value := (*s)[n]
	*s = (*s)[0:n]
	return value
}
