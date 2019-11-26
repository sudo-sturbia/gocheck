// Load words from a file into a Trie to be used as a dictionary
package loader

import (
	"bufio"
	"log"
	"os"
)

// A trie node
type Node struct {
	Children [SIZE_OF_ALPHABET]*Node // An array of nodes
	IsWord   bool                    // True if node references a word ending, false otherwise
}

const SIZE_OF_ALPHABET = 26

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
	if root.Children[word[charNumber]-'a'] == nil {
		root.Children[word[charNumber]-'a'] = new(Node)
	}

	root.Children[word[charNumber]-'a'] = loadWord(root.Children[word[charNumber]-'a'], word, charNumber+1)
	return root
}
