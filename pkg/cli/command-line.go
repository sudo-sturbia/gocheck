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
		"By default a word that contains\nan uppercase letter any where but the start is considered\nwrong," +
		"when this flag is used, this feature is disabled.\n"

	uppercaseShort := flag.Bool("u", false, uString)
	uppercaseLong := flag.Bool("uppercase", false, uString)

	// Ignore word flag
	iString := "ignore specified word (specified word is considered correct.) " +
		"This\nflag can be used an unlimited amount of times.\n"

	flag.Var(new(ignore), "i", iString)
	flag.Var(new(ignore), "ignore", iString)

	// Help flag
	helpShort := flag.Bool("h", false, "Print this help message.")
	helpLong := flag.Bool("help", false, "Print this help message.")

	// Parsing
	flag.Parse()

	if *helpShort || *helpLong {
		fmt.Println(
			"gocheck is a simple, fast spell-checker.\n" +
				"It works by comparing a file against a given list of words and printing errors.\n" +
				"\n" +
				"Usage\n" +
				"\n" +
				"    gocheck [OPTIONS] <FILEPATH> <DICTIONARYPATH>\n" +
				"\n" +
				"Required arguments:\n" +
				"\n" +
				"    <FILEPATH>        Path to a text file that should be processed to find errors.\n" +
				"    <DICTIONARYPATH>  Path to a text file containing a list of words, one word per\n" +
				"                      line, to compare the other file against.\n" +
				"\n" +
				"Options:\n" +
				"\n" +
				"    -h --help         Print this help message.\n" +
				"\n" +
				"    -i --ignore WORD  Ignore specified word (WORD is considered correct.) This\n" +
				"                      flag can be used an unlimited amount of times.\n" +
				"\n" +
				"    -u --uppercase    Ignore uppercase letters. By default a word that contains\n" +
				"                      an uppercase letter any where but the start is considered\n" +
				"                      wrong, when this flag is used, this feature is disabled.\n" +
				"\n" +
				"For the source code check the github page [github.com/sudo-sturbia/gocheck]\n",
		)
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
