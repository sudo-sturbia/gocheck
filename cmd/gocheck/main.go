package main

import (
    "fmt"
    "os"
    "bufio"
)

const SIZE_OF_ALPHABET = 26
const FIRST_ASCII_CHAR = 97

func main() {
    loadDictionary()
}

// A trie node
type Node struct {
    childNodes [SIZE_OF_ALPHABET]*Node    // An array of child nodes
    isWord bool                           // True if node references a word ending, false otherwise
}

// Load a dictionay of words into a Trie
func loadDictionary() *Node {
    root := new(Node)

    // Open words' file
    file, err := os.Open("words.txt")

    if err != nil {
        fmt.Println(err)
        panic(err)
    }

    defer file.Close()

    // Load words from file to a trie
    fileScanner := bufio.NewScanner(file)
    for fileScanner.Scan() {
        word := fileScanner.Text()

        root = loadWord(root, word, 0)
    }

    fmt.Println(root.childNodes)

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
    if root.childNodes[word[charNumber] - FIRST_ASCII_CHAR] == nil {
        root.childNodes[word[charNumber] - FIRST_ASCII_CHAR] = new(Node)
    }

    root.childNodes[word[charNumber] - FIRST_ASCII_CHAR] = loadWord(root.childNodes[word[charNumber] - FIRST_ASCII_CHAR], word, charNumber + 1)
    return root
}
