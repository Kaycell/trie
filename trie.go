package trie

import "strings"

type node struct {
	content    rune
	wordMarker bool
	children   []*node
}

func (n *node) findChild(r rune) *node {
	for _, e := range n.children {
		if e.content == r {
			return e
		}
	}
	return nil
}

// Trie represents a trie also called prefix tree.
// The zero value for Trie is an empty trie ready to use.
type Trie struct {
	root   node
	nbNode int
}

// Init initializes or clears trie t.
func (t *Trie) Init() *Trie {
	t.root.content = 0
	t.root.children = nil
	t.root.wordMarker = false
	t.nbNode = 0
	return t
}

// New returns an initialized trie.
func New() *Trie { return new(Trie).Init() }

// AddWord adds word in t.
// Word is stored one rune per node in t.
// If the root of the word is already present only missing runes are added.
// wordmarker is set to true on the node that contains last word rune.
func (t *Trie) AddWord(word string) {
	currentNode := &t.root

	if len(word) == 0 {
		currentNode.wordMarker = true
	}

	for i, c := range word {
		child := currentNode.findChild(c)

		if child == nil {
			newNode := &node{content: c}
			currentNode.children = append(currentNode.children, newNode)
			t.nbNode++
			currentNode = newNode
		} else {
			currentNode = child
		}

		if i == len(word)-1 {
			currentNode.wordMarker = true
		}
	}
}

// SearchWord searches a path in t that contains every word runes and
// end by wordmarked node.
func (t *Trie) SearchWord(word string) bool {
	currentNode := &t.root

	for _, c := range word {
		child := currentNode.findChild(c)
		if child == nil {
			return false
		}
		currentNode = child
	}
	return currentNode.wordMarker
}

// findAllWords use DFS algorithm to find all words from node n
func findAllWords(n *node, prefix string, words []string) []string {
	if n.wordMarker {
		words = append(words, prefix)
	}
	for _, child := range n.children {
		prefix = prefix + string(child.content)
		words = findAllWords(child, prefix, words)
		prefix = strings.TrimSuffix(prefix, string(child.content))
	}
	return words
}

// FindAllWords returns all words present in the trie
func (t *Trie) FindAllWords() []string {
	var words []string
	return findAllWords(&t.root, "", words)
}

// FindAllMatchingWords returns all words present in the trie that
// match the prefix
func (t *Trie) FindAllMatchingWords(prefix string) []string {
	var matchs []string
	currentNode := &t.root
	for _, c := range prefix {
		child := currentNode.findChild(c)
		if child == nil {
			return matchs
		}
		currentNode = child
	}
	matchs = findAllWords(currentNode, prefix, matchs)
	return matchs
}
