// Package main initializes application and calls other packages.
package main

import (
	"github.com/sudo-sturbia/gocheck/pkg/checker"
	"github.com/sudo-sturbia/gocheck/pkg/loader"
)

// Initialize program and parse command line flags
func main() {
	spellChecker := checker.Instance()

	// Find paths
	// TODO ..
	filePath := ""
	dictionaryPath := ""

	dictionary := loader.LoadDictionary(dictionaryPath)

	spellChecker.CheckFile(dictionary, filePath)
	spellChecker.PrintSpellingErrors()
}
