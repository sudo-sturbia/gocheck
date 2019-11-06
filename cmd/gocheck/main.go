package main

import (
    "bufio"
    "errors"
    "flag"
    "fmt"
    "log"
    "os"
    "strings"
    "unicode"
)

const SIZE_OF_ALPHABET = 26
const FIRST_ASCII_CHAR = 97

func main() {
    // Process command line arguments
    filePath := flag.String("f", "", "Path to the file that should be processed.")
    dictionaryPath := flag.String("d", "", "Path to dictionary used in processing.")

    filePathError := errors.New("option -f empty, no file specified.")
    dictionaryPathError := errors.New("option -d empty, no dictionary file specified.")

    flag.Parse()

    if len(*filePath) == 0 {
        log.Fatal(filePathError)
    } else if len(*dictionaryPath) == 0 {
        log.Fatal(dictionaryPathError)
    }

    dictionary := loadDictionary(*dictionaryPath)
    checkFile(dictionary, *filePath)
}

/*
 * Loading a dictionary
 */

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
    if root.children[word[charNumber] - FIRST_ASCII_CHAR] == nil {
        root.children[word[charNumber] - FIRST_ASCII_CHAR] = new(Node)
    }

    root.children[word[charNumber] - FIRST_ASCII_CHAR] = loadWord(root.children[word[charNumber] - FIRST_ASCII_CHAR], word, charNumber + 1)
    return root
}

/*
 * Verifying file
 */

// Check file for spelling errors
func checkFile(root *Node, path string) {
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

        words := strings.FieldsFunc(textLine, wordEnd)

        for i := 0; i < len(words); i++ {
            if !checkWord(root, strings.ToLower(words[i]), 0) {
                // Print spelling error
                fmt.Printf("Error at (line: %d, word: %d)  \"%s\" incorrect.\n", lineNumber, i, words[i])
            }
        }

        lineNumber++
    }

    if err := fileScanner.Err(); err != nil {
        log.Fatal(err)
    }
}

// Check if a word exists in the dictionary
func checkWord(root *Node, word string, charNumber int) bool {
    // If end of word
    if charNumber == len(word) {
        return root.isWord
    }

    if root.children[word[charNumber] - FIRST_ASCII_CHAR] != nil {
        return checkWord(root.children[word[charNumber] - FIRST_ASCII_CHAR], word, charNumber + 1)
    } else {
        return false
    }
}
