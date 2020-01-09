// Package loader implements functions for loading words
// from a file into a trie to be used as a dictionary.
package loader

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

const (
	printableASCII      = 95
	firstPrintableASCII = 32
)

// Node represents a node in a trie.
type Node struct {
	children [printableASCII]*Node // Children nodes
	isWord   bool                  // True if node marks a word ending, false otherwise
}

// Children returns an array of children of a Node.
func (n *Node) Children() [printableASCII]*Node {
	return n.children
}

// IsWord returns true if node marks a word
// ending, false otherwise.
func (n *Node) IsWord() bool {
	return n.isWord
}

// LoadDictionary loads a dictionay of words into a trie,
// and returns a pointer to trie's head Node.
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

// Load a word into trie.
// Return a pointer to trie's head node.
func loadWord(root *Node, word string, charNumber int) *Node {
	// If end of word
	if charNumber == len(word) {
		root.isWord = true
		return root
	}

	// If Node is not initialized
	if root.children[word[charNumber]-firstPrintableASCII] == nil {
		root.children[word[charNumber]-firstPrintableASCII] = new(Node)
	}

	// If character is printable ASCII
	if word[charNumber] >= 32 && word[charNumber] <= unicode.MaxASCII {
		root.children[word[charNumber]-firstPrintableASCII] = loadWord(root.children[word[charNumber]-firstPrintableASCII], word, charNumber+1)
	}

	return root
}
