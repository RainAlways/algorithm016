package other

import (
	"log"
	"testing"
)

func TestWordDictionary(t *testing.T) {
	wordDictionary := WConstructor()
	wordDictionary.AddWord("bad")
	log.Println(wordDictionary.Search("bad"))
}
