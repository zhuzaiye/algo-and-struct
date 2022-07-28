package strings

import (
	"fmt"
	"testing"
)

func TestStringTrie(t *testing.T) {
	trie := NewTrie()
	words := []string{"Golang", "程序员", `Go!*$\n`, "Lang"}
	for _, word := range words {
		trie.Insert(word)
	}
	tar := `Go!*$\n`
	if trie.Find(tar) {
		fmt.Printf("Find string %s\n", tar)
	} else {
		fmt.Printf("Not Find string %s\n", tar)
	}
}
