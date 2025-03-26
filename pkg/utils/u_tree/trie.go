package u_tree

// TrieNode 是前缀树的节点
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool // 标记是否为单词的结尾
}

// Trie 是前缀树
type Trie struct {
	root *TrieNode
}

// NewTrie 创建一个新的前缀树
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{children: make(map[rune]*TrieNode)}}
}

// Insert 向前缀树中插入一个单词
func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			node.children[ch] = &TrieNode{children: make(map[rune]*TrieNode)}
		}
		node = node.children[ch]
	}
	node.isEnd = true
}

// Search 查找前缀树中是否存在某个单词
func (t *Trie) Search(word string) bool {
	node := t.root
	for _, ch := range word {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return node.isEnd
}

// StartsWith 查找是否存在以某个前缀开头的单词
func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, ch := range prefix {
		if _, exists := node.children[ch]; !exists {
			return false
		}
		node = node.children[ch]
	}
	return true
}
