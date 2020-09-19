package other

func findWords(board [][]byte, words []string) []string {
	strs := make([]string, 0, 5)
	trie := buildTrie(words)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dns(board, i, j, trie.root, &strs)
		}
	}

	return strs
}

func dns(board [][]byte, row, col int, node *TrieNode2, strs *[]string) {
	c := board[row][col]
	if c == '#' || node.children[int(c-'a')] == nil {
		return
	}

	node = node.children[int(c-'a')]
	if node.word != "" {
		*strs = append(*strs, node.word)
		node.word = ""
	}

	board[row][col] = '#'

	if row > 0 {
		dns(board, row-1, col, node, strs)
	}
	if col > 0 {
		dns(board, row, col-1, node, strs)
	}
	if row < len(board)-1 {
		dns(board, row+1, col, node, strs)
	}
	if col < len(board[row])-1 {
		dns(board, row, col+1, node, strs)
	}
	board[row][col] = c
}

func buildTrie(words []string) *Trie2 {
	t := &Trie2{root: initTrieNode2()}

	for _, word := range words {
		t.insertWord(word)
	}

	return t
}

type Trie2 struct {
	root *TrieNode2
}

type TrieNode2 struct {
	children [26]*TrieNode2
	word     string
}

func initTrieNode2() *TrieNode2 {
	return &TrieNode2{
		children: [26]*TrieNode2{},
	}
}

func (t Trie2) insertWord(word string) {
	w := t.root
	for _, c := range word {
		index := int(c - 'a')
		if w.children[index] == nil {
			w.children[index] = initTrieNode2()
		}
		w = w.children[index]
	}

	w.word = word
}
