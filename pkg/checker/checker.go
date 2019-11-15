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
var IgnoreUppercase bool

var mux sync.Mutex
var Wg sync.WaitGroup

const SIZE_OF_ALPHABET = 26

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
func checkLine(root *loader.Node, textLine string, lineNumber int, wordEnd func (c rune) bool) {
    defer Wg.Done()

    words := strings.FieldsFunc(textLine, wordEnd)

    for i := 0; i < len(words); i++ {
        // Ignore case if flag is used
        if IgnoreUppercase {
            words[i] = strings.ToLower(words[i])
        }

        if !checkWord(root, words[i], 0) {
            // Add spelling error to list
            mux.Lock()
            spellingErrors = append(spellingErrors, fmt.Sprintf("At (%d, %d)  \"%s\"", lineNumber, i, words[i]))
            mux.Unlock()
        }
    }
}

// Check if a word exists in the dictionary
func checkWord(root *loader.Node, word string, charNumber int) bool {

    if charNumber == len(word) {
        return root.IsWord
    }

    // Check type of character
    if unicode.IsLetter(rune(word[charNumber])) {

        if unicode.IsUpper(rune(word[charNumber])) {
            if charNumber == 0 {
                if root.Children[byte(unicode.ToLower(rune(word[charNumber]))) - 'a'] != nil {
                    return checkWord(root.Children[byte(unicode.ToLower(rune(word[charNumber]))) - 'a'], word, charNumber + 1)
                }
            }

            return false
        } else {
            // Check if character exists
            if root.Children[word[charNumber] - 'a'] != nil {
                return checkWord(root.Children[word[charNumber] - 'a'], word, charNumber + 1)
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
