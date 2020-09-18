package other

type Trie struct {
	root *TrieNode
}

type TrieNode struct {
	children [26]*TrieNode
	isWord   bool
}

/** Initialize your data structure here. */
func TConstructor() Trie {
	return Trie{
		root: newTrieNode(),
	}
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: [26]*TrieNode{},
		isWord:   false,
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	w := this.root
	for _, value := range word {
		index := int(value - 'a')
		if w.children[index] == nil {
			w.children[index] = newTrieNode()
		}
		w = w.children[index]
	}
	w.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	w := this.root
	for _, value := range word {
		index := int(value - 'a')
		if w.children[index] == nil {
			return false
		}
		w = w.children[index]
	}

	return w.isWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	w := this.root
	for _, value := range prefix {
		index := int(value - 'a')
		if w.children[index] == nil {
			return false
		}
		w = w.children[index]
	}

	return true
}
