# gocheck

> A simple, fast spell-checker.

## How it works

gocheck is a spell-checker that works by comparing words in a text file against a list of words, both are given as arguments.

The list is loaded into memory creating a dictionary. Words from the file are then compared one by one against the dictionary, specifying positions of incorrect ones and printing error messages accordingly.

## How to install

```
go get github.com/sudo-sturbia/gocheck/cmd/gocheck
```

## How to use

```console
usage:

       gocheck [OPTIONS] <FILEPATH> <DICTIONARYPATH>

required arguments:

       <FILEPATH>        Path to a text file that should be processed to find errors.
       <DICTIONARYPATH>  Path to a text file containing a list of words, one word per line, to compare the the other file against.

options:

       -h --help         Print this help message.
       -i WORD           Ignore specified word (word is considered correct.) This flag can be used an unlimited amount of times.
       -u                Ignore uppercase letters.
                         By default a word that contains an uppercase letter any where but the start is considered wrong, when this flag is used, this feature is disabled.
```

A comprehensive dictionary can be downloaded from [english-words](https://github.com/dwyl/english-words/blob/master/words_alpha.txt).

#### Usage example

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

## Implementation

gocheck uses several methods to enhance processing time,

- The given list of words is loaded from a text file into a **Trie**.
- Text lines are processed concurrently using **goroutines**.
- An error message is specified for each incorrect word, messages are collectively printed after processing.

