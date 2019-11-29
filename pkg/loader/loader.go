// Load words from a file into a Trie to be used as a dictionary
package loader

import (
	"bufio"
	"log"
	"os"
	"unicode"
)

// A trie node
type Node struct {
	Children [PRINTABLE_ASCII]*Node // An array of nodes
	IsWord   bool                   // True if node references a word ending, false otherwise
}

const PRINTABLE_ASCII = 95
const FIRST_PRINTABLE_ASCII = 32

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
		root.IsWord = true
		return root
	}

	// If Node is not initialized
	if root.Children[word[charNumber]-FIRST_PRINTABLE_ASCII] == nil {
		root.Children[word[charNumber]-FIRST_PRINTABLE_ASCII] = new(Node)
	}

	// If character is printable ASCII
	if word[charNumber] >= 32 && word[charNumber] <= unicode.MaxASCII {
		root.Children[word[charNumber]-FIRST_PRINTABLE_ASCII] = loadWord(root.Children[word[charNumber]-FIRST_PRINTABLE_ASCII], word, charNumber+1)
	}

	return root
}
