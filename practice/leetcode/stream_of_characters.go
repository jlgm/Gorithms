type trieNode struct {
	childrens [26]*trieNode
	isWordEnd bool
}

type trie struct {
	root *trieNode
}

func initTrie() *trie {
	return &trie{
		root: &trieNode{},
	}
}

func (t *trie) insert(word *string) {
	wordLen := len(*word)
	current := t.root
	for i := wordLen - 1; i >= 0; i-- {
		index := (*word)[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &trieNode{}
		}
		current = current.childrens[index]
	}
	current.isWordEnd = true
}

func (t *trie) find(word *[]byte) bool {
	wordLen := len(*word)
	cur := t.root
	for i := wordLen - 1; i >= 0; i-- {
		idx := (*word)[i] - 'a'
		if cur.childrens[idx] == nil {
			return false
		} else if cur.childrens[idx].isWordEnd {
			return true
		}
		cur = cur.childrens[idx]
	}
	return cur.isWordEnd
}

type StreamChecker struct {
	stream []byte
	trie   *trie
}

func Constructor(words []string) StreamChecker {
	stream := []byte("")

	trie := initTrie()
	for _, w := range words {
		trie.insert(&w)
	}

	return StreamChecker{trie: trie, stream: stream}
}

func (this *StreamChecker) Query(letter byte) bool {
	this.stream = append(this.stream, letter)
	return this.trie.find(&this.stream)
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */