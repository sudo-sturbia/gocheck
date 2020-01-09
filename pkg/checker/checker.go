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

const (
	PRINTABLE_ASCII       = 95
	FIRST_PRINTABLE_ASCII = 32
)

// Concurrency related variables.
var (
	mux sync.Mutex
	wg  sync.WaitGroup
)

// Checker, holds data related to
// spelling errors and verification.
type Checker struct {
	spellingErrors []string // A list of spelling errors

	ignoredWords    map[string]bool // A map of words to ignore
	ignoreUppercase bool            // Consider all given words to be lowercase
}

// New returns pointer to a new, initialized Checker object.
func New() *Checker {
	instance := new(Checker)
	instance.ignoredWords = make(map[string]bool)

	return instance
}

// Add a word to the ignored words list.
func (c *Checker) AddWordToIgnored(word string) {
	c.ignoredWords[word] = true
}

// Add a given list of words to ignored words.
func (c *Checker) AddListToIgnored(words []string) {
	for _, word := range words {
		c.ignoredWords[word] = true
	}
}

// Return a string containing all ignored words.
func (c *Checker) IgnoredString() string {
	var stringForm string
	for key, _ := range c.ignoredWords {
		stringForm += fmt.Sprintf("%s ", key)
	}

	return stringForm
}

// Set ignore uppercase boolean.
func (c *Checker) SetIgnoreUppercase(ignore bool) {
	c.ignoreUppercase = true
}

// Check file for spelling errors.
func (c *Checker) CheckFile(root *loader.Node, path string) {
	c.spellingErrors = make([]string, 0)

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

	spellingErrorsInLine := make([]string, 0)
	words := strings.FieldsFunc(textLine, wordEnd)

	hasErrors := false

	for i, word := range words {
		if !c.ignoredWords[word] {
			if c.ignoreUppercase {
				word = strings.ToLower(word)
			}

			if !c.checkWord(root, word, 0) {
				// Add formatted error to list
				spellingErrorsInLine = append(spellingErrorsInLine, fmt.Sprintf("At (%d, %d)  \"%s\"", lineNumber, i, word))
				hasErrors = true
			}
		}
	}

	// Add line's errors to errors slice
	if hasErrors {
		mux.Lock()
		c.spellingErrors = append(c.spellingErrors, spellingErrorsInLine...)
		mux.Unlock()
	}
}

// Return true if a word exists in the trie,
// return false otherwise.
func (c *Checker) checkWord(root *loader.Node, word string, charNumber int) bool {

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
					return c.checkWord(root.Children()[word[charNumber]], word, charNumber+1)
				}
			}

			return false
		} else {
			// Check if character exists
			if root.Children()[word[charNumber]-FIRST_PRINTABLE_ASCII] != nil {
				return c.checkWord(root.Children()[word[charNumber]-FIRST_PRINTABLE_ASCII], word, charNumber+1)
			}

			return false
		}
	} else {
		return false
	}
}

// Print spelling errors.
func (c *Checker) PrintSpellingErrors() {
	for _, spellingError := range c.spellingErrors {
		fmt.Println(spellingError)
	}

	fmt.Printf("- Found a total of %d errors.\n", len(c.spellingErrors))
}
