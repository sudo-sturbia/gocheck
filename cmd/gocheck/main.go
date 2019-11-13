package main

import (
    "bufio"
    "errors"
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
    "sync"
    "unicode"
)

var spellingErrors []string
var ignoreUppercase bool

var mux sync.Mutex
var wg sync.WaitGroup

func main() {
    // Process command line flags
    filePathFlag := flag.String("f", "", "Path to the file that should be processed.")
    dictionaryPathFlag := flag.String("d", "", "Path to dictionary used in processing.")
    uppercaseFlag := flag.Bool("u", false, "Ignore uppercase letters. By default a word that contains an uppercase letter any where but the start is considered wrong, " +
                                           "when this flag is used uppercase and lowercase letters are treated similarly.")

    flag.Parse()

    // If no path is specified
    if len(*filePathFlag) == 0 {
        log.Fatal(errors.New("option -f empty, no file specified."))
    } else if len(*dictionaryPathFlag) == 0 {
        log.Fatal(errors.New("option -d empty, no dictionary file specified."))
    }

    ignoreUppercase = *uppercaseFlag

    dictionary := loadDictionary(*dictionaryPathFlag)
    checkFile(dictionary, *filePathFlag)

    wg.Wait()
    printSpellingErrors()
}

/*
 * Loading a dictionary
 */

const SIZE_OF_ALPHABET = 26

// A trie node
type Node struct {
    children [SIZE_OF_ALPHABET]*Node      // An array of nodes
    isWord bool                           // True if node references a word ending, false otherwise
}

// Load a dictionay of words into a Trie
func loadDictionary(path string) *Node {
    // Open dictionary file
    file, err := os.Open(path)

    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    // Load words from file to a trie
    root := new(Node)

    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        word := fileScanner.Text()

        root = loadWord(root, word, 0)
    }

    if err := fileScanner.Err(); err != nil {
        log.Fatal(err)
    }

    return root
}

// Load a word into Trie
func loadWord(root *Node, word string, charNumber int) *Node {
    // If end of word
    if charNumber == len(word) {
        root.isWord = true
        return root
    }

    // If Node is not initialized
    if root.children[word[charNumber] - 'a'] == nil {
        root.children[word[charNumber] - 'a'] = new(Node)
    }

    root.children[word[charNumber] - 'a'] = loadWord(root.children[word[charNumber] - 'a'], word, charNumber + 1)
    return root
}

/*
 * Verifying file
 */

// Check file for spelling errors
func checkFile(root *Node, path string) {
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

        wg.Add(1)
        go checkLine(root, textLine, lineNumber, wordEnd)

        lineNumber++
    }

    if err := fileScanner.Err(); err != nil {
        log.Fatal(err)
    }
}

// Find spelling errors in a line
func checkLine(root *Node, textLine string, lineNumber int, wordEnd func (c rune) bool) {
    defer wg.Done()

    words := strings.FieldsFunc(textLine, wordEnd)

    for i := 0; i < len(words); i++ {
        // Ignore case if flag is used
        if ignoreUppercase {
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
func checkWord(root *Node, word string, charNumber int) bool {

    if charNumber == len(word) {
        return root.isWord
    }

    // Check type of character
    if unicode.IsLetter(rune(word[charNumber])) {

        if unicode.IsUpper(rune(word[charNumber])) {
            if charNumber == 0 {
                if root.children[byte(unicode.ToLower(rune(word[charNumber]))) - 'a'] != nil {
                    return checkWord(root.children[byte(unicode.ToLower(rune(word[charNumber]))) - 'a'], word, charNumber + 1)
                }
            }

            return false
        } else {
            // Check if character exists
            if root.children[word[charNumber] - 'a'] != nil {
                return checkWord(root.children[word[charNumber] - 'a'], word, charNumber + 1)
            }

            return false
        }
    } else {
        return false
    }
}

// Print spelling errors
func printSpellingErrors() {
    numberOfErrors := len(spellingErrors)

    for i := 0; i < numberOfErrors; i++ {
        fmt.Println(spellingErrors[i])
    }

    fmt.Printf("- Found a total of %d errors.\n", numberOfErrors)
}
