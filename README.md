# gocheck

> A simple, fast spell-checker.

## How To Install

```
go get github.com/sudo-sturbia/gocheck/cmd/gocheck
```

## How To Use
### Package
For package documentation see [checker](https://pkg.go.dev/github.com/sudo-sturbia/gocheck/pkg/checker),
and [loader](https://pkg.go.dev/github.com/sudo-sturbia/gocheck/pkg/loader).

```go
// Usage Example
// Load a Trie to use as a dictionary
dictionary := loader.LoadList([]string{
	"list",
	"of",
	"words",
	"to",
	"verify",
	"against",
})

// Set of words to verify
words := []string{
	"wrds",
	"against",
}

c := checker.New()
errors := c.CheckList(dictionary, words)
for _, word := range errors {
	fmt.Println(word)
}

// Output: wrds
```

### Command Line Tool
```console
gocheck is a simple, fast spell-checker.
It works by comparing a file against a given list of words and prints
spelling errors accordingly.

Usage
    gocheck [options] <filepath> <dictionarypath>

Required Arguments
    <filepath>        Path to a text file to spellcheck.
    <dictionarypath>  Path to a text file containing a list of words, one word per
                      line, to spellcheck against.

Options
    -h              Print a short help message.
    -help           Print a detailed help message.
    -ignore <word>  Ignore given word (consider it correct.)
    -ignore-upper   By default a word that contains an uppercase letter any where
                    but the start is considered wrong. When this flag is used, this
                    behaviour is disabled.

For the source code see [github.com/sudo-sturbia/gocheck]
```

#### Example
Using files in test-data/ directory

```
gocheck test-data/wrong-paragraph.txt test-data/test-words.txt
```

Output

```console
At (3, 2)  "nevsdfser"
At (3, 9)  "rmation"
At (2, 11)  "th"
At (1, 2)  "s12eleted"
At (1, 4)  "stu"
At (1, 5)  "ck"
At (0, 3)  "memmorable"
At (0, 9)  "mde"
- Found a total of 8 errors.
```

A comprehensive dictionary can be downloaded from [english-words](https://github.com/dwyl/english-words).
