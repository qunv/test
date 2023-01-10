package main

type Node struct {
	child     [255]*Node
	isWordEnd bool
	lines     []int
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

func (t *Trie) Insert(word string, line int) {
	wordLength := len(word)
	pivot := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if pivot.child[index] == nil {
			pivot.child[index] = &Node{}
		}
		pivot = pivot.child[index]
	}
	pivot.isWordEnd = true
	pivot.lines = append(pivot.lines, line)
}

func (t *Trie) Search(word string) []int {
	pivot := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if pivot.child[index] == nil {
			return nil
		}
		pivot = pivot.child[index]
	}
	if pivot.isWordEnd {
		return pivot.lines
	}
	return nil
}
