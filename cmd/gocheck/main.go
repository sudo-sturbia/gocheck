// Package main initializes application, calls other packages,
// and handles parsing of command line arguments, and flags.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sudo-sturbia/gocheck/v3/pkg/checker"
	"github.com/sudo-sturbia/gocheck/v3/pkg/loader"
)

// ignored is a list of words to ignore while spell-checking. -ignore
// flag is used to specify words.
type ignored []string

// Command line flags
var (
	ignoredWords = make(ignored, 16)
	shortH       = flag.Bool("h", false, "Print a short help message.")
	detailedH    = flag.Bool("help", false, "Print a detailed help message.")
	upper        = flag.Bool("ignore-upper", false, "By default a word that contains an uppercase letter any where "+
		"but the start is considered wrong. When this flag is used, this behaviour is disabled.")
)

func main() {
	filePath, dictionaryPath := parse()

	dictionary := loader.LoadFile(dictionaryPath)

	c := checker.New()
	c.IgnoreList(ignoredWords)
	c.SetIgnoreUppercase(*upper)

	errors, err := c.CheckFile(dictionary, filePath)
	if err != nil {
		log.Fatal(err)
	}

	for _, word := range errors {
		fmt.Printf("At (%d, %d) \"%s\"\n", word.Row, word.Col, word.Word)
	}

	fmt.Printf("- Found a total of %d errors.\n", len(errors))
}

// parse parses command line arguments and flags. Returns two paths,
// a file to verify, and a dictionary file.
func parse() (string, string) {
	flag.Var(&ignoredWords, "ignore", "Ignore given word (consider it correct.)")
	flag.Parse()

	help()

	file := flag.Arg(0)
	dictionary := flag.Arg(1)
	if file == "" || dictionary == "" { // Print short help message
		usage()
		os.Exit(0)
	}

	return file, dictionary
}

// help prints a short or a detailed help message, if -h or -help
// are respectively used.
func help() {
	if *shortH {
		usage()
		os.Exit(0)
	}

	if *detailedH {
		man()
		os.Exit(0)
	}
}

// usage displays a short usage message.
func usage() {
	fmt.Printf(
		"Usage\n" +
			"\tgocheck [options] <filepath> <dictionarypath>\n" +
			"Use -help for more details.\n")
}

// man displays a help message if either -h, or -help are used.
func man() {
	fmt.Printf(
		"gocheck is a simple, fast spell-checker.\n" +
			"\n" +
			"It works by comparing a file against a given list of words and prints\n" +
			"spelling errors accordingly.\n" +
			"\n" +
			"Usage\n" +
			"\tgocheck [options] <filepath> <dictionarypath>\n" +
			"\n" +
			"Required Arguments\n" +
			"\t<filepath>        Path to a text file to spellcheck.\n" +
			"\t<dictionarypath>  Path to a text file containing a list of words, one word per\n" +
			"\t                  line, to spellcheck against.\n" +
			"\n" +
			"Options\n" +
			"\t-h              Print a short help message.\n" +
			"\t-help           Print a detailed help message.\n" +
			"\t-ignore <word>  Ignore given word (consider it correct.)\n" +
			"\t-ignore-upper   By default a word that contains an uppercase letter any where\n" +
			"\t                but the start is considered wrong. When this flag is used, this\n" +
			"\t                behaviour is disabled.\n" +
			"\n" +
			"For the source code see [github.com/sudo-sturbia/gocheck]\n")
}

// String returns string representation.
func (i *ignored) String() string {
	builder := new(strings.Builder)
	for _, s := range *i {
		builder.WriteString(s)
		builder.WriteByte(' ')
	}

	return builder.String()
}

// Set value adds given string to list of ignored strings.
func (i *ignored) Set(value string) error {
	if value != "" {
		ignoredWords = append(ignoredWords, value)
	}
	return nil
}
