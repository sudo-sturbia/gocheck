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

// Test file checking on a file without errors.
func TestCheckFileWithoutErrors(t *testing.T) {
	testChecker := New()
	testChecker.CheckFile(root, "../../test-data/paragraph.txt")

	if len(testChecker.errors) != 0 {
		t.Errorf("Number of spelling errors incorrect, expected: 0, got: %d", len(testChecker.errors))
	}
}

// Test file checking on a file with errors.
func TestCheckFileWithErrors(t *testing.T) {
	testChecker := New()
	testChecker.CheckFile(root, "../../test-data/wrong-paragraph.txt")

	if len(testChecker.errors) != 8 {
		t.Errorf("Number of spelling errors incorrect, expected: 8, got: %d", len(testChecker.errors))
	}

	// Assert errors found in file
	// Create error map
	errorMap := make(map[string]bool)

	errorMap["At (0, 3)  \"memmorable\""] = true
	errorMap["At (0, 9)  \"mde\""] = true
	errorMap["At (1, 2)  \"s12eleted\""] = true
	errorMap["At (1, 4)  \"stu\""] = true
	errorMap["At (1, 5)  \"ck\""] = true
	errorMap["At (2, 11)  \"th\""] = true
	errorMap["At (3, 2)  \"nevsdfser\""] = true
	errorMap["At (3, 9)  \"rmation\""] = true

	for i := 0; i < 8; i++ {
		if !errorMap[testChecker.errors[i]] {
			t.Errorf("Found incorrect error: %s", testChecker.errors[i])
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

// Benchmark processing time.
func BenchmarkWordProcessing(b *testing.B) {
	testChecker := New()
	for n := 0; n < b.N; n++ {
		testChecker.CheckFile(root, "../../test-data/paragraph.txt")
	}
}
