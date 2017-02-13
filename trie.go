package trie

type node struct {
	content    rune
	children   []*node
	wordMarker bool
}

func (n *node) findChild(r rune) *node {
	for _, e := range n.children {
		if e.content == r {
			return e
		}
	}
	return nil
}

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

func (t *Trie) SearchWord(word string) bool {
	currentNode := &t.root

	for _, c := range word {
		child := currentNode.findChild(c)
		if child == nil {
			return false
		}
		currentNode = child
	}
	if currentNode.wordMarker {
		return true
	}
	return false
}
