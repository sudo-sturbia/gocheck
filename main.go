package main

const SIZE_OF_ALPHABET = 26

func main() {
    // ...
}

// A trie node
type node struct {
    childNodes [SIZE_OF_ALPHABET]*node    // An array of child nodes
    isWord bool                           // True if node references a word ending, false otherwise
}

// Load a dictionay of words into a trie
// Returns an array of nodes
func loadDictionary() [SIZE_OF_ALPHABET]*node {
    // ...
}
