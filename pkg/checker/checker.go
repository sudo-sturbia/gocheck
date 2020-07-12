// Package checker implements functions used to find spelling errors
// in a given text file and print error messages accordingly.
package checker

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"unicode"

	"github.com/sudo-sturbia/gocheck/pkg/loader"
)

// Concurrency related variables.
var (
	mux sync.Mutex
	wg  sync.WaitGroup
)

// Checker holds data related to
// spelling errors and verification.
type Checker struct {
	errors          []string        // List of spelling errors
	ignored         map[string]bool // Map of words to ignore
	ignoreUppercase bool            // Consider all given words to be lowercase
}

// New returns pointer to a new, initialized Checker object.
func New() *Checker {
	instance := new(Checker)
	instance.ignored = make(map[string]bool)

	return instance
}

// Ignore adds a word to the ignored words list.
func (c *Checker) Ignore(word string) {
	c.ignored[word] = true
}

// IgnoreList adds a given list of words to ignored words.
func (c *Checker) IgnoreList(words []string) {
	for _, word := range words {
		c.ignored[word] = true
	}
}

// SetIgnoreUppercase sets Checker's ignoreUppercase flag.
func (c *Checker) SetIgnoreUppercase(ignore bool) {
	c.ignoreUppercase = true
}

// CheckFile checks file for spelling errors and populates
// Checker's errors list.
func (c *Checker) CheckFile(root *loader.Node, path string) {
	// Open file
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordEnd := func(c rune) bool {
		return unicode.IsPunct(c) || (c == ' ')
	}

	// Read words from file
	lineNumber := 0

	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		textLine := fileScanner.Text()

		wg.Add(1)
		go c.checkLine(root, textLine, lineNumber, wordEnd)

		lineNumber++
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}

// Finds spelling errors in a line (string) of words.
// Adds found errors to errors array.
// Spelling errors are formatted as --- At (row, word)  "Error".
func (c *Checker) checkLine(root *loader.Node, textLine string, lineNumber int, wordEnd func(c rune) bool) {
	defer wg.Done()

	lineErrors := make([]string, 0)
	words := strings.FieldsFunc(textLine, wordEnd)

	hasErrors := false

	for i, word := range words {
		if !c.ignored[word] {
			if c.ignoreUppercase {
				word = strings.ToLower(word)
			}

			if !c.CheckWord(root, word, 0) {
				// Add formatted error to list
				lineErrors = append(lineErrors, fmt.Sprintf("At (%d, %d)  \"%s\"", lineNumber, i, word))
				hasErrors = true
			}
		}
	}

	// Add line's errors to errors slice
	if hasErrors {
		mux.Lock()
		c.errors = append(c.errors, lineErrors...)
		mux.Unlock()
	}
}

// Return true if a word exists in the trie,
// return false otherwise.
func (c *Checker) CheckWord(root *loader.Node, word string, charNumber int) bool {

	if charNumber == len(word) {
		return root.IsWord()
	}

	// Check if character is ASCII
	if word[charNumber] >= 32 && word[charNumber] <= unicode.MaxASCII {

		// Uppercase character
		if word[charNumber] >= 65 && word[charNumber] <= 90 {
			if charNumber == 0 {
				// Pass character as uppercase
				if root.Children()[word[charNumber]] != nil {
					return c.CheckWord(root.Children()[word[charNumber]], word, charNumber+1)
				}
			}

			return false
		}

		// Check if character exists
		if root.Children()[word[charNumber]-loader.FirstPrintableASCII] != nil {
			return c.CheckWord(root.Children()[word[charNumber]-loader.FirstPrintableASCII], word, charNumber+1)
		}

		return false

	}

	return false
}

// PrintSpellingErrors prints strings in Checker's spellingErrors list.
func (c *Checker) PrintSpellingErrors() {
	for _, spellingError := range c.errors {
		fmt.Println(spellingError)
	}

	fmt.Printf("- Found a total of %d errors.\n", len(c.errors))
}
