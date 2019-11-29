// Find spelling errors in a text file and print error messages
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

var spellingErrors []string
var IgnoredWords map[string]bool
var IgnoreUppercase bool

var mux sync.Mutex
var Wg sync.WaitGroup

const PRINTABLE_ASCII = 95
const FIRST_PRINTABLE_ASCII = 32

// Check file for spelling errors
func CheckFile(root *loader.Node, path string) {
	spellingErrors = make([]string, 0)

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

		Wg.Add(1)
		go checkLine(root, textLine, lineNumber, wordEnd)

		lineNumber++
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Find spelling errors in a line
func checkLine(root *loader.Node, textLine string, lineNumber int, wordEnd func(c rune) bool) {
	defer Wg.Done()

	spellingErrorsInLine := make([]string, 0)
	words := strings.FieldsFunc(textLine, wordEnd)

	hasErrors := false
	for i := 0; i < len(words); i++ {
		if !IgnoredWords[words[i]] {
			if IgnoreUppercase {
				words[i] = strings.ToLower(words[i])
			}

			if !checkWord(root, words[i], 0) {
				// Add spelling error to list
				spellingErrorsInLine = append(spellingErrorsInLine, fmt.Sprintf("At (%d, %d)  \"%s\"", lineNumber, i, words[i]))
				hasErrors = true
			}
		}
	}

	// Add line's errors to errors' slice
	if hasErrors {
		mux.Lock()
		spellingErrors = append(spellingErrors, spellingErrorsInLine...)
		mux.Unlock()
	}
}

// Check if a word exists in the dictionary
func checkWord(root *loader.Node, word string, charNumber int) bool {

	if charNumber == len(word) {
		return root.IsWord
	}

	// Check if character is ASCII
	if word[charNumber] >= 32 && word[charNumber] <= unicode.MaxASCII {

		// Uppercase character
		if word[charNumber] >= 65 && word[charNumber] <= 90 {
			if charNumber == 0 {
				// Pass character as uppercase
				if root.Children[word[charNumber]] != nil {
					return checkWord(root.Children[word[charNumber]], word, charNumber+1)
				}
			}

			return false
		} else {
			// Check if character exists
			if root.Children[word[charNumber]-FIRST_PRINTABLE_ASCII] != nil {
				return checkWord(root.Children[word[charNumber]-FIRST_PRINTABLE_ASCII], word, charNumber+1)
			}

			return false
		}
	} else {
		return false
	}
}

// Print spelling errors
func PrintSpellingErrors() {
	numberOfErrors := len(spellingErrors)

	for i := 0; i < numberOfErrors; i++ {
		fmt.Println(spellingErrors[i])
	}

	fmt.Printf("- Found a total of %d errors.\n", numberOfErrors)
}
