// Package main initializes application, calls other packages,
// and handles parsing of command line arguments, and flags.
package main

import (
	"github.com/sudo-sturbia/gocheck/pkg/checker"
	"github.com/sudo-sturbia/gocheck/pkg/loader"
)

var spellChecker *checker.Checker

// Initialize program and parse command line flags
func main() {
	spellChecker = checker.New()

	// Get paths
	filePath, dictionaryPath := parse()

	dictionary := loader.LoadDictionary(dictionaryPath)

	spellChecker.CheckFile(dictionary, filePath)
	spellChecker.PrintSpellingErrors()
}
