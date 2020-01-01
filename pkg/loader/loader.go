// Load words from a file into a Trie to be used as a dictionary
package loader

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

const (
	PRINTABLE_ASCII       = 95
	FIRST_PRINTABLE_ASCII = 32
)

// A trie node
type Node struct {
	children [PRINTABLE_ASCII]*Node // Children nodes
	isWord   bool                   // True if node marks a word ending, false otherwise
}

// Return an array of children nodes.
func (n *Node) Children() [PRINTABLE_ASCII]*Node {
	return n.children
}

// Returns true if node marks a word
// ending, false otherwise.
func (n *Node) IsWord() bool {
	return n.isWord
}

// Load a dictionay of words into a Trie
func LoadDictionary(path string) *Node {
	// Open dictionary file
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Load words from file to a trie
	root := new(Node)

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		word := fileScanner.Text()

		root = loadWord(root, word, 0)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return root
}

// Load a word into Trie
func loadWord(root *Node, word string, charNumber int) *Node {
	// If end of word
	if charNumber == len(word) {
		root.isWord = true
		return root
	}

	// If Node is not initialized
	if root.children[word[charNumber]-FIRST_PRINTABLE_ASCII] == nil {
		root.children[word[charNumber]-FIRST_PRINTABLE_ASCII] = new(Node)
	}

	// If character is printable ASCII
	if word[charNumber] >= 32 && word[charNumber] <= unicode.MaxASCII {
		root.children[word[charNumber]-FIRST_PRINTABLE_ASCII] = loadWord(root.children[word[charNumber]-FIRST_PRINTABLE_ASCII], word, charNumber+1)
	}

	return root
}
