package checker

import (
	"testing"

	"github.com/sudo-sturbia/gocheck/pkg/loader"
)

// Root of used trie.
var root = loader.LoadList([]string{
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
})

// Test CheckList function on a list without errors
func TestCheckListWithoutErrors(t *testing.T) {
	c := New()
	words := []string{
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

	if len(c.CheckList(root, words)) != 0 {
		t.Errorf("Found error in a list without errors.")
	}
}

// Test CheckList function on a list containing errors
func TestCheckListWithErrors(t *testing.T) {
	c := New()
	words := []string{
		"read",
		"ths",
		"moment",
		"long",
		"chi",
		"ion",
		"o",
		"gd",
		"thorns",
		"_-",
		"th",
		"never",
		"bound",
		"but",
		"formation",
		"first",
		"lik",
		"on",
	}

	shouldFind := []string{
		"ths",
		"chi",
		"ion",
		"o",
		"gd",
		"_-",
		"th",
		"lik",
	}

	found := c.CheckList(root, words)
	if len(found) != len(shouldFind) {
		t.Errorf("Expected %d errors, found %d.", len(shouldFind), len(found))
	}

	// Push found errors into a map
	foundMap := make(map[string]bool)
	for _, word := range found {
		foundMap[word] = true
	}

	// Compare found with shouldFind
	for _, word := range shouldFind {
		if !foundMap[word] {
			t.Errorf("Didn't find %s.", word)
		}
	}
}

// Test file checking on a file without errors.
func TestCheckFileWithoutErrors(t *testing.T) {
	c := New()
	found, err := c.CheckFile(root, "../../test-data/paragraph.txt")
	if err != nil {
		t.Errorf("File reading failed.")
	}

	if len(found) != 0 {
		t.Errorf("Found errors in a correct file.")
	}
}

// Test file checking on a file with errors.
func TestCheckFileWithErrors(t *testing.T) {
	c := New()
	found, err := c.CheckFile(root, "../../test-data/wrong-paragraph.txt")
	if err != nil {
		t.Errorf("File reading failed.")
	}

	if len(found) != 8 {
		t.Errorf("Incorrect number of spelling errors. Expected 8, Found %d.", len(found))
	}

	// Assert errors found in file
	shouldFind := []SpellingError{
		SpellingError{"memmorable", 0, 3},
		SpellingError{"mde", 0, 9},
		SpellingError{"s12eleted", 1, 2},
		SpellingError{"stu", 1, 4},
		SpellingError{"ck", 1, 5},
		SpellingError{"th", 2, 11},
		SpellingError{"nevsdfser", 3, 2},
		SpellingError{"rmation", 3, 9},
	}

	// Push found errors to a map
	foundMap := make(map[string]SpellingError)
	for _, err := range found {
		foundMap[err.Word] = err
	}

	// Compare found with shouldFind
	for _, err := range shouldFind {
		if foundMap[err.Word] != err {
			t.Errorf("Didn't find %v.", err)
		}
	}
}

// Test word checking using correct words.
func TestCheckWordExists(t *testing.T) {
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
	}

	for _, word := range words {
		if !CheckWord(root, word) {
			t.Errorf("\"%s\" should exist, but doesn't.", word)
		}
	}
}

// Test word checking using incorrect words.
func TestCheckWordDoesntExist(t *testing.T) {
	words := []string{
		"df",
		"ad",
		"thhhink",
		"howsds",
		"ifferent",
		"ts",
		"curse",
		"iwould",
		"hve",
		"beeen",
		"pse",
		"yu",
		"wh^o",
		"$$$",
		"\\",
		"monet",
		"lll",
		"chan",
	}

	for _, word := range words {
		if CheckWord(root, word) {
			t.Errorf("\"%s\" shouldn't exist, but does.", word)
		}
	}
}

// Benchmark CheckList function.
func BenchmarkCheckList(b *testing.B) {
	c := New()
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

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.CheckList(root, words)
	}
}

// Benchmark CheckFile function.
func BenchmarkCheckFile(b *testing.B) {
	c := New()
	for n := 0; n < b.N; n++ {
		c.CheckFile(root, "../../test-data/paragraph.txt")
	}
}
