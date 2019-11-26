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
	filePathFlag := flag.String("f", "", "path to the file that should be processed.")
	dictionaryPathFlag := flag.String("d", "", "path to dictionary used validation.")
	uppercaseFlag := flag.Bool("u", false, "ignore uppercase letters. By default a word that contains an uppercase letter any where but the start is considered wrong, "+
		"when this flag is used, this feature is disabled.")
	flag.Var(new(ignoreFlag), "i", "ignore specified word (specified word is considered correct.) This flag can be used an unlimited amount of times.")
	flag.Usage = helpFlag

	flag.Parse()

	// If no path is specified
	if len(*filePathFlag) == 0 {
		log.Fatal(errors.New("option -f empty, no file specified."))
	} else if len(*dictionaryPathFlag) == 0 {
		log.Fatal(errors.New("option -d empty, no dictionary file specified."))
	}

	checker.IgnoreUppercase = *uppercaseFlag

	dictionary := loader.LoadDictionary(*dictionaryPathFlag)
	checker.CheckFile(dictionary, *filePathFlag)

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
		"gocheck is a simple, fast spell-checker.\n" +
			"\n" +
			"usage:\n" +
			"\n" +
			"       gocheck [-h] [-f PATH] [-d PATH] [-i WORD] [-u]\n" +
			"\n" +
			"required arguments:\n" +
			"\n" +
			"       -f PATH     Path to the file that should be processed.\n" +
			"       -d PATH     Path to dictionary used for validation.\n" +
			"                   A dictionary is a file containing a collection of lowercase words, one word per line.\n" +
			"\n" +
			"optional arguments:\n" +
			"\n" +
			"       -h --help   Print this help message.\n" +
			"       -i WORD     Ignore WORD (specified word is considered correct.) This flag can be used an unlimited amount of times.\n" +
			"       -u          Ignore uppercase letters.\n" +
			"                   By default a word that contains an uppercase letter any where but the start is considered wrong, when this flag is used, this feature is disabled.\n" +
			"\n" +
			"For the source code check the github page [github.com/sudo-sturbia/gocheck]",
	)
}
