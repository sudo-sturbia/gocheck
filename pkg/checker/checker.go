// Package checker implements a simple, fast spell-checker.
package checker

import (
	"bufio"
	"log"
	"os"
	"strings"
	"unicode"

	"github.com/sudo-sturbia/gocheck/pkg/loader"
)

// Checker is used to find spelling errors. Checker implements
// several options when spell-checking such as ignored words, and
// words with uppercase letters.
type Checker struct {
	ignored         map[string]bool // Map of words to ignore
	ignoreUppercase bool            // Consider all given words to be lowercase
}

// SpellingError represents a spelling error found in a text file.
type SpellingError struct {
	Word string // Incorrectly spelled word.
	Row  int    // Row containing the word.
	Col  int    // Column containing the word.
}

// New returns pointer to a new, initialized Checker object.
func New() *Checker {
	return &Checker{make(map[string]bool), false}
}

// Ignore adds a word to ignored words.
func (c *Checker) Ignore(word string) {
	c.ignored[word] = true
}

// IgnoreList adds a given list of words to ignored words.
func (c *Checker) IgnoreList(words []string) {
	for _, word := range words {
		c.ignored[word] = true
	}
}

// ClearIgnored clears Checker's ignored words list.
func (c *Checker) ClearIgnored(ignored bool) {
	if ignored {
		for i := range c.ignored {
			delete(c.ignored, i)
		}
	}
}

// SetIgnoreUppercase sets Checker's ignoreUppercase flag. By default
// a word with an uppercase letter anywhere but the start is considered
// wrong. When ignoreUppercase is true, this behaviour is disabled.
func (c *Checker) SetIgnoreUppercase(ignore bool) {
	c.ignoreUppercase = ignore
}

// CheckList checks a list of strings against a given Trie and returns
// a slice containing incorrect words.
func (c *Checker) CheckList(root *loader.Node, list []string) []string {
	counter := make(chan bool) // To count checked words.
	errorChan := make(chan string)
	for _, word := range list {
		if c.ignoreUppercase {
			word = strings.ToLower(word)
		}

		go func(word string) {
			if !c.ignored[word] && !CheckWord(root, word) {
				errorChan <- word
			}

			counter <- true
		}(word)
	}

	go func() {
		for count := 0; count < len(list); count++ {
			<-counter
		}

		close(errorChan)
	}()

	errors := make([]string, 0)
	for word := range errorChan {
		errors = append(errors, word)
	}

	return errors
}

// CheckFile checks the file at given path for spelling errors against
// a given Trie. Returns a list of incorrect words with their row and
// column numbers, and an IO error if file reading fails.
func (c *Checker) CheckFile(root *loader.Node, path string) ([]SpellingError, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	errorChan := make(chan SpellingError)
	done := make(chan bool)

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		go c.CheckLine(root, scanner.Text(), errorChan, done, line, func(c rune) bool {
			return unicode.IsPunct(c) || (c == ' ')
		})

		line++
	}

	// If an IO error occurred
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	go func() {
		for count := 0; count < line; count++ {
			<-done
		}

		close(errorChan)
	}()

	errors := make([]SpellingError, 0)
	for word := range errorChan {
		errors = append(errors, word)
	}

	return errors, nil
}

// CheckLine takes a line of text (string containing multiple words), seperates the
// line into words using wordEnd function, checks each word in the line against
// the given trie, and pushes incorrect words to errorChan. After line evaluation is
// finished, true is sent as a singal to done channel.
func (c *Checker) CheckLine(root *loader.Node, line string, errorChan chan SpellingError, done chan bool, lineNumber int, wordEnd func(c rune) bool) {
	wordDone := make(chan bool)
	words := strings.FieldsFunc(line, wordEnd)
	for i, word := range words {
		if c.ignoreUppercase {
			word = strings.ToLower(word)
		}

		go func(word string, index int) {
			if !c.ignored[word] && !CheckWord(root, word) {
				errorChan <- SpellingError{word, lineNumber, index}
			}

			wordDone <- true
		}(word, i)
	}

	for count, length := 0, len(words); count < length; count++ {
		<-wordDone
	}

	done <- true
}

// CheckWord verifies a given word against a given Trie, returns
// true if word exists in the given trie, false otherwise.
func CheckWord(root *loader.Node, word string) bool {
	return recCheck(root, word, 0)
}

// recCheck, recursively, verifies that a word exists in the given trie.
func recCheck(root *loader.Node, word string, charNumber int) bool {
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
					return recCheck(root.Children()[word[charNumber]], word, charNumber+1)
				}
			}

			return false
		}

		// Check if character exists
		if root.Children()[word[charNumber]-loader.FirstPrintableASCII] != nil {
			return recCheck(root.Children()[word[charNumber]-loader.FirstPrintableASCII], word, charNumber+1)
		}

		return false

	}

	return false
}
