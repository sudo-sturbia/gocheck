package main

import (
    "errors"
    "flag"
    "fmt"
    "log"

    "github.com/sudo-sturbia/gocheck/pkg/loader"
    "github.com/sudo-sturbia/gocheck/pkg/checker"
)

func main() {
    // Process command line flags
    filePathFlag := flag.String("f", "", "path to the file that should be processed.")
    dictionaryPathFlag := flag.String("d", "", "Path to dictionary used validation.")
    uppercaseFlag := flag.Bool("u", false, "ignore uppercase letters. By default a word that contains an uppercase letter any where but the start is considered wrong, " +
                                           "when this flag is used, this feature is disabled.")

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

// Create --help flag
func helpFlag() {
    fmt.Println(
        "gocheck is a simple, fast spell-checker.\n" +
        "\n" +
        "usage:\n" +
        "\n" +
        "       gocheck [-h] [-f PATH] [-d PATH] [-u]\n" +
        "\n" +
        "required arguments:\n" +
        "\n" +
        "       -f PATH     Path to the file that should be processed.\n" +
        "       -d PATH     Path to dictionary used for validation.\n" +
        "\n" +
        "optional arguments:\n" +
        "\n" +
        "       -h --help   Print this help message.\n" +
        "       -u          Ignore uppercase letters.\n" +
        "                   By default a word that contains an uppercase letter any where but the start is considered wrong, when this flag is used, this feature is disabled.\n" +
        "\n" +
        "For the source code check the github page [github.com/sudo-sturbia/gocheck]",
    )
}
