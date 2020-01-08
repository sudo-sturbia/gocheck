# gocheck

> A simple, fast spell-checker.

## How To Install

```
go get github.com/sudo-sturbia/gocheck/cmd/gocheck
```

## How To Use

```console
gocheck is a simple, fast spell-checker.
It works by comparing a file against a given list of words and printing errors.

Usage

    gocheck [OPTIONS] <FILEPATH> <DICTIONARYPATH>

Required arguments:

    <FILEPATH>        Path to a text file that should be processed to find errors.
    <DICTIONARYPATH>  Path to a text file containing a list of words, one word per
                      line, to compare the other file against.

Options:

    -h --help         Print this help message.

    -i --ignore WORD  Ignore specified word (WORD is considered correct.) This
                      flag can be used an unlimited amount of times.

    -u --uppercase    Ignore uppercase letters. By default a word that contains
                      an uppercase letter any where but the start is considered
                      wrong, when this flag is used, this feature is disabled.

For the source code check the github page [github.com/sudo-sturbia/gocheck]
```

### Example

Using files in test/ directory

```
gocheck test/paragraph-wrong.txt test/test_words.txt
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

## How It Works
gocheck works by comparing a given file against another containing a list of words.

Incorrect words are collected and error messages are printed accordingly.
Error messages are formatted as `At (%lineNumber, %wordNumber) "Word"`.

gocheck is optimized for performance:

- a trie data structure is used to load dictionary.
- goroutines are used to process input files.

