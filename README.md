# gocheck

> A simple, fast spell-checker.

## How To Install

```
go get github.com/sudo-sturbia/gocheck/cmd/gocheck
```

## How To Use

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

### Example
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

## How It Works
gocheck compares a given text file against another containing a list of
words. Incorrect words are collected, and an error message formatted as
```
At (%LineNumber, %WordNumber) "incorrectWord"
.
.
- Found a total of N errors.
```
is printed.

gocheck optimizes perfomance by using

- Trie data structure for dictionary loading/querying,
- Goroutines to process input files.
