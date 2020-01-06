// Package cli handles parsing of command line
// arguments and flags.
package cli

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/sudo-sturbia/gocheck/pkg/checker"
)

// Parse command line arguments and flags.
// Return paths to file to verify and dictionary.
func Parse() (string, string) {
	// Ignore uppercase flag
	uString := "ignore uppercase letters. " +
		"By default a word that contains an uppercase letter any where but the start is considered wrong, \n" +
		"when this flag is used, this feature is disabled."

	uppercaseShort := flag.Bool("u", false, uString)
	uppercaseLong := flag.Bool("uppercase", false, uString)

	// Ignore word flag
	iString := "ignore specified word (specified word is considered correct.) " +
		"This flag can be used an unlimited amount of times."

	flag.Var(new(ignore), "i", iString)
	flag.Var(new(ignore), "ignore", iString)

	// Help flag
	helpShort := flag.Bool("h", false, "")
	helpLong := flag.Bool("help", false, "")

	// Parsing
	flag.Parse()

	if *helpShort || *helpLong {
		// TODO
		fmt.Println("")
		os.Exit(0)
	}

	checker.Instance().SetIgnoreUppercase(*uppercaseShort || *uppercaseLong)

	// Command line arguments
	file := flag.Arg(0)
	dictionary := flag.Arg(1)

	// If no path is specified
	if file == "" {
		log.Fatal(errors.New("no file specified."))
	} else if dictionary == "" {
		log.Fatal(errors.New("no dictionary file specified."))
	}

	return file, dictionary
}

// ignore flag is used to specify words to ignore when spell checking.
type ignore string

// Return string representation.
func (i *ignore) String() string {
	return "Ignored words: " + checker.Instance().IgnoredString()
}

// Set value adds given string to list of ignored strings.
func (i *ignore) Set(value string) error {
	checker.Instance().AddIgnoredWord(value)
	return nil
}
