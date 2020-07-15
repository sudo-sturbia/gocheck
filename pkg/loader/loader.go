// Package loader implements functions for loading strings into a trie
// to be used as a dictionary.
package loader

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

// Number of printable ASCII characters and their starting position.
const (
	PrintableASCII      = 95
	FirstPrintableASCII = 32
)

// Node represents a node in a trie.
type Node struct {
	children [PrintableASCII]*Node // Children nodes
	isWord   bool                  // True if node marks a word ending, false otherwise
}

// Children returns an array of children of a Node.
func (n *Node) Children() [PrintableASCII]*Node {
	return n.children
}

// IsWord returns true if node marks a word ending, false otherwise.
func (n *Node) IsWord() bool {
	return n.isWord
}

// LoadFile creates a new trie using words from text file at the
// given path. Returns a pointer to trie's root.
func LoadFile(path string) *Node {
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

		root = LoadWord(root, word)
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	return root
}

// LoadList creates a new trie using given string list. Returns a pointer
// to trie's root node.
func LoadList(list []string) *Node {
	root := new(Node)
	for _, word := range list {
		root = LoadWord(root, word)
	}

	return root
}

// LoadWord loads a word into trie. Returns a pointer to trie's root
// node.
func LoadWord(root *Node, word string) *Node {
	return recLoad(root, word, 0)
}

// recLoad loads a word into a trie recursively.
func recLoad(root *Node, word string, whichChar int) *Node {
	// If end of word
	if whichChar == len(word) {
		root.isWord = true
		return root
	}

	// If Node is not initialized
	if root.children[word[whichChar]-FirstPrintableASCII] == nil {
		root.children[word[whichChar]-FirstPrintableASCII] = new(Node)
	}

	// If character is printable ASCII
	if word[whichChar] >= 32 && word[whichChar] <= unicode.MaxASCII {
		root.children[word[whichChar]-FirstPrintableASCII] = recLoad(root.children[word[whichChar]-FirstPrintableASCII], word, whichChar+1)
	}

	return root
}
