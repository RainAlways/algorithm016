package Week_02

import (
	"log"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	log.Println(groupAnagrams(strs))
}
