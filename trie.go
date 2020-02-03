package trie

import "errors"

// Trie struct.
type Trie struct {
	rootNode  *Node // root node
	wordCount int   // total dictionary
}

// Node struct.
type Node struct {
	prefix         rune           // prefix
	parent         *Node          // parent node
	childrens      map[rune]*Node // children nodes
	IsCompleteWord bool           // is complete word
	ResultCount    int            // total same prefix nodes.
	Payload        interface{}
}

// New create a new trie
func New() *Trie {
	return &Trie{
		rootNode: &Node{
			prefix:         0,
			parent:         nil,
			childrens:      make(map[rune]*Node),
			IsCompleteWord: false,
			ResultCount:    0,
		},
		wordCount: 0,
	}
}

// Insert new word into trie.
// adding new nodes as needed.
func (t *Trie) Insert(word string, payload interface{}) {
	// root node from trie
	currentNode := t.rootNode

	for _, character := range word {
		// change current node if there was exists character in current node children
		if node, found := currentNode.childrens[character]; found {
			currentNode = node
			node.ResultCount++
		} else {
			// the character is not exists in current node,
			// create new node children.
			currentNode.childrens[character] = &Node{
				prefix:         character,
				parent:         currentNode,
				childrens:      make(map[rune]*Node),
				IsCompleteWord: false,
				ResultCount:    1,
			}

			// set current node to created node children.
			currentNode = currentNode.childrens[character]
		}
	}

	// because we loop all character of the word,
	// so the currentNode is complete word.
	currentNode.IsCompleteWord = true
	currentNode.Payload = payload
	// increment total words/dictionary in trie.
	t.wordCount++
}

// Find will return the matched node.
func (t *Trie) Find(word string) (*Node, error) {
	currentNode := t.rootNode

	var err error

	for _, character := range word {
		if node, found := currentNode.childrens[character]; found {
			currentNode = node
		} else {
			// there is no matching node.
			err = errors.New("no matching node")
			break
		}
	}

	return currentNode, err
}

// Remove will delete nodes.
func (t *Trie) Remove(word string) bool {
	node, err := t.Find(word)

	if err != nil {
		return false
	}

	if !node.IsCompleteWord {
		return false
	}

	for node.parent != nil {
		// delete all children nodes.
		node.ResultCount--
		if node.ResultCount == 0 {
			delete(node.parent.childrens, node.prefix)
		}

		// scan parent node.
		node = node.parent
	}

	t.wordCount--
	return true
}
