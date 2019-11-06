package main

import (
    "fmt"
    "os"
    "bufio"
    "flag"
    "errors"
    "log"
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

    loadDictionary(*dictionaryPath)
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

    fmt.Println(root.children)

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
