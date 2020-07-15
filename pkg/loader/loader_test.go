package loader

import (
	"testing"
)

// Test loading of one word.
func TestWordLoading(t *testing.T) {
	root := new(Node)
	LoadWord(root, "word")

	if !isWordLoaded(root, "word", 0) {
		t.Errorf("Word \"word\" was not loaded.\n")
	}
}

// Test loading from a file.
func TestLoadingFromFile(t *testing.T) {
	root := LoadFile("../../test-data/test-load.txt")

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

	for _, word := range words {
		if !isWordLoaded(root, word, 0) {
			t.Errorf("Word \"%s\" was not loaded.\n", word)
		}
	}
}

// Test loading a trie from a list.
func TestLoadingFromList(t *testing.T) {
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

	root := LoadList(words)
	for _, word := range words {
		if !isWordLoaded(root, word, 0) {
			t.Errorf("Word \"%s\" was not loaded.\n", word)
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

// Benchmark loading from a file.
func BenchmarkLoadingFromFile(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LoadFile("../../test-data/test-words.txt")
	}
}

// Benchmark loading from a list.
func BenchmarkLoadingFromList(b *testing.B) {
	words := []string{
		"that",
		"was",
		"a",
		"memorable",
		"day",
		"to",
		"me",
		"for",
		"it",
		"made",
		"great",
		"changes",
		"in",
		"but",
		"is",
		"the",
		"same",
		"with",
		"any",
		"life",
		"imagine",
		"one",
		"selected",
		"struck",
		"out",
		"of",
		"and",
		"think",
		"how",
		"different",
		"its",
		"course",
		"would",
		"have",
		"been",
		"pause",
		"you",
		"who",
		"read",
		"this",
		"moment",
		"long",
		"chain",
		"iron",
		"or",
		"gold",
		"thorns",
		"flowers",
		"that",
		"never",
		"bound",
		"but",
		"formation",
		"first",
		"link",
		"on",
	}

	for n := 0; n < b.N; n++ {
		LoadList(words)
	}
}
