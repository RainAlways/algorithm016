package other

import (
	"log"
	"testing"
)

func TestFindWords(t *testing.T) {
	var bytes [][]byte
	bytes = append(bytes, []byte{
		'o', 'a', 'a', 'n',
	})
	bytes = append(bytes, []byte{
		'e', 't', 'a', 'e',
	})
	bytes = append(bytes, []byte{
		'i', 'h', 'k', 'r',
	})
	bytes = append(bytes, []byte{
		'i', 'f', 'l', 'v',
	})
	words := []string{"oath", "pea", "eat", "rain"}
	log.Println(findWords(bytes, words))
}
