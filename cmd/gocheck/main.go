package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/sudo-sturbia/gocheck/pkg/checker"
	"github.com/sudo-sturbia/gocheck/pkg/loader"
)

// Initialize program and parse command line flags
func main() {
	// Process command line flags
	uppercaseFlag := flag.Bool("u", false, "ignore uppercase letters. By default a word that contains an uppercase letter any where but the start is considered wrong, \n"+
		"when this flag is used, this feature is disabled.")
	flag.Var(new(ignoreFlag), "i", "ignore specified word (specified word is considered correct.) This flag can be used an unlimited amount of times.")
	flag.Usage = helpFlag

	flag.Parse()

	// Get arguments
	filePath := flag.Arg(0)
	dictionaryPath := flag.Arg(1)

	// If no path is specified
	if filePath == "" {
		log.Fatal(errors.New("no file specified."))
	} else if dictionaryPath == "" {
		log.Fatal(errors.New("no dictionary file specified."))
	}

	checker.IgnoreUppercase = *uppercaseFlag

	dictionary := loader.LoadDictionary(dictionaryPath)
	checker.CheckFile(dictionary, filePath)

	checker.Wg.Wait()
	checker.PrintSpellingErrors()
}

// Create custom ignore flag implementing flag.Value interface
type ignoreFlag string

// Return string representation
func (i *ignoreFlag) String() string {
	var stringRepresentation string
	for key := range checker.IgnoredWords {
		stringRepresentation += key
	}

	return stringRepresentation
}

// Set flag value
func (i *ignoreFlag) Set(value string) error {
	if checker.IgnoredWords == nil {
		checker.IgnoredWords = make(map[string]bool)
	}

	checker.IgnoredWords[value] = true
	return nil
}

// Create --help flag
func helpFlag() {
	fmt.Println(
		"gocheck is a simple, fast spell-checker. It compares a text file against a list of given words and finds incorrect ones.\n" +
			"\n" +
			"usage:\n" +
			"\n" +
			"       gocheck [OPTIONS] <FILEPATH> <DICTIONARYPATH>\n" +
			"\n" +
			"required arguments:\n" +
			"\n" +
			"       <FILEPATH>        Path to a text file that should be processed to find errors.\n" +
			"       <DICTIONARYPATH>  Path to a text file containing a list of words, one word per line, to compare the the other file against.\n" +
			"\n" +
			"options:\n" +
			"\n" +
			"       -h --help         Print this help message.\n" +
			"       -i WORD           Ignore specified word (word is considered correct.) This flag can be used an unlimited amount of times.\n" +
			"       -u                Ignore uppercase letters.\n" +
			"                         By default a word that contains an uppercase letter any where but the start is considered wrong, when this flag is used, this feature is disabled.\n" +
			"\n" +
			"For the source code check the github page [github.com/sudo-sturbia/gocheck]",
	)
}
