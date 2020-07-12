// Package loader implements functions for loading words
// from a file into a trie to be used as a dictionary.
package loader

import (
	"testing"
)

// Test loading of one word.
func TestWordLoading(t *testing.T) {
	root := new(Node)
	LoadWord(root, "word", 0)

	if !isWordLoaded(root, "word", 0) {
		t.Errorf("Word \"word\" was not loaded.\n")
	}
}

// Test dictionary loading.
func TestDictionaryLoading(t *testing.T) {
	root := LoadDictionary("../../test/test-load.txt")

	// Test loaded words
	words := []string{
		"this",
		"is",
		"a",
		"simple",
		"list",
		"used",
		"to",
		"test",
		"loading",
	}

	for i := 0; i < len(words); i++ {
		if !isWordLoaded(root, words[i], 0) {
			t.Errorf("Word \"%s\" was not loaded.\n", words[i])
		}
	}
}

// Return true if word is correctly loaded, false otherwise.
func isWordLoaded(root *Node, word string, whichChar int) bool {
	if whichChar == len(word) {
		return root.isWord
	}

	return isWordLoaded(root.children[word[whichChar]-FirstPrintableASCII], word, whichChar+1)
}

// Benchmark dictionary loading.
func BenchmarkDictionaryLoading(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LoadDictionary("../../test/test-words.txt")
	}
}
