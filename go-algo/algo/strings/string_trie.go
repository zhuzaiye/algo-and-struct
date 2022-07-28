package strings

// Trie树，又称为字典树，用于处理字符串匹配，用于解决字符串的匹配查找
// 整体复杂度 O(n*m)

// 根据字符串集合构造Trie树

// Trie树节点定义
type trieNode struct {
	char     string             // Unicode字符
	isEnding bool               // 是否是单词结尾
	children map[rune]*trieNode // 该节点的子节点字典
}

// 初始化Trie树
func NewTrieNode(char string) *trieNode {
	return &trieNode{
		char:     char,
		isEnding: false,
		children: make(map[rune]*trieNode),
	}
}

// Trie
type Trie struct {
	root *trieNode
}

// 初始化Trie
func NewTrie() *Trie {
	trieNode := NewTrieNode("/") // 创建以“/”的初始化树
	return &Trie{
		root: trieNode,
	}
}

// 将字符串插入树
func (t *Trie) Insert(word string) {
	node := t.root // 根节点
	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			value = NewTrieNode(string(code))
			node.children[code] = value
		}
		node = value
	}
	node.isEnding = true
}

// 在Trie中查找一个单词
func (t *Trie) Find(word string) bool {
	node := t.root
	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			return false
		}
		node = value
	}
	if !node.isEnding {
		return false // 没有完全匹配
	}
	return true
}
