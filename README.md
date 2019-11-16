# gocheck
> A simple, fast spell-checker.

## How to install

```
go get github.com/sudo-sturbia/gocheck/cmd/gocheck
```

## How to use

```console
usage:

       gocheck [-h] [-f PATH] [-d PATH] [-i WORD] [-u]

required arguments:

       -f PATH     Path to the file that should be processed.
       -d PATH     Path to dictionary used for validation.
                   A dictionary is a file containing a collection of lowercase words, one word per line.

optional arguments:

       -h --help   Print this help message.
       -i WORD     Ignore WORD (specified word is considered correct.) This flag can be used an unlimited amount of times.
       -u          Ignore uppercase letters.
                   By default a word that contains an uppercase letter any where but the start is considered wrong, when this flag is used, this feature is disabled.
```

A dictionary can be downloaded from [english-words](https://github.com/dwyl/english-words/blob/master/words_alpha.txt).

#### Usage example

Using files in test/ directory

```
gocheck -f test/paragraph-wrong.txt -d test/test_words.txt
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

## How it works

gocheck uses several methods to enhance processing time,

- Words are loaded from a text file, specified by `-d` flag, into a **Trie** to be used as a dictionary for word verification.
- Text lines from the file, specified by `-f` flag, are processed concurrently using **goroutines**.
- For each incorrect word a message is specified, error messages are collectively printed after processing.
