package Week_02

import (
	"log"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	log.Println(topKFrequent(nums, k))
}
