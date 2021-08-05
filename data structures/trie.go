type trieNode struct {
	childrens [26]*trieNode
	isWordEnd bool
}

type Trie struct {
	root *trieNode
}

/** Initialize your data structure here. */
func NewTrie() Trie {
	return Trie{
		root: &trieNode{},
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	cur := this.root
	for _, c := range word {
		idx := (c - 'a')
		if cur.childrens[idx] == nil {
			cur.childrens[idx] = &trieNode{}
		}
		cur = cur.childrens[idx]
	}
	cur.isWordEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	cur := this.root
	for _, c := range word {
		idx := (c - 'a')
		if cur.childrens[idx] == nil {
			return false
		}
		cur = cur.childrens[idx]
	}
	return cur.isWordEnd
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix {
		idx := (c - 'a')
		if cur.childrens[idx] == nil {
			return false
		}
		cur = cur.childrens[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := NewTrie();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */