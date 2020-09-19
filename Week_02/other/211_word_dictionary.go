package other

type WordDictionary struct {
	root *WordNode
}

type WordNode struct {
	children [26]*WordNode
	isWord   bool
}

func initWordNode() *WordNode {
	return &WordNode{
		children: [26]*WordNode{},
	}
}

/** Initialize your data structure here. */
func WConstructor() WordDictionary {
	return WordDictionary{root: initWordNode()}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	w := this.root
	for _, value := range word {
		index := int(value - 'a')
		if w.children[index] == nil {
			w.children[index] = initWordNode()
		}
		w = w.children[index]
	}

	w.isWord = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return matchWord([]byte(word), 0, this.root)
}

func matchWord(bytes []byte, pos int, node *WordNode) bool {
	if node == nil {
		return false
	}

	if len(bytes) == pos {
		return node.isWord
	}

	char := bytes[pos]
	if char != '.' {
		return matchWord(bytes, pos+1, node.children[int(char-'a')])
	} else {
		for i := 0; i < 26; i++ {
			if matchWord(bytes, pos+1, node.children[i]) {
				return true
			}
		}
	}

	return false
}
