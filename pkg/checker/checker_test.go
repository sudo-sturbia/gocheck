// Package checker implements functions used to find spelling errors
// in a given text file and print error messages accordingly.
package checker

import (
	"os"
	"testing"

	"github.com/sudo-sturbia/gocheck/pkg/loader"
)

var (
	testChecker *checker
	root        *loader.Node
)

// Setup dictionary before testing.
func TestMain(m *testing.M) {
	root = loader.LoadDictionary(os.Getenv("GOPATH") + "/src/github.com/sudo-sturbia/gocheck/test/test_words.txt")
	os.Exit(m.Run())
}

// Test check file function on a file without errors.
func TestCheckFileWithoutErrors(t *testing.T) {
	testChecker = reInit()
	testChecker.CheckFile(root, os.Getenv("GOPATH")+"/src/github.com/sudo-sturbia/gocheck/test/paragraph.txt")

	if len(testChecker.spellingErrors) != 0 {
		t.Errorf("Number of spelling errors incorrect, expected: 0, got: %d", len(testChecker.spellingErrors))
	}
}

// Test check file function on a file with errors.
func TestCheckFileWithErrors(t *testing.T) {
	testChecker = reInit()
	testChecker.CheckFile(root, os.Getenv("GOPATH")+"/src/github.com/sudo-sturbia/gocheck/test/paragraph-wrong.txt")

	if len(testChecker.spellingErrors) != 8 {
		t.Errorf("Number of spelling errors incorrect, expected: 8, got: %d", len(testChecker.spellingErrors))
	}

	// Assert errors found in file
	// Create error map
	errorMap := make(map[string]bool)

	errorMap["At (0, 3)  \"memmorable\""] = true
	errorMap["At (0, 9)  \"mde\""] = true
	errorMap["At (1, 2)  \"s12eleted\""] = true
	errorMap["At (1, 4)  \"stu\""] = true
	errorMap["At (1, 5)  \"ck\""] = true
	errorMap["At (2, 11)  \"th\""] = true
	errorMap["At (3, 2)  \"nevsdfser\""] = true
	errorMap["At (3, 9)  \"rmation\""] = true

	for i := 0; i < 8; i++ {
		if !errorMap[testChecker.spellingErrors[i]] {
			t.Errorf("Found incorrect error: %s", testChecker.spellingErrors[i])
		}
	}

}

// Benchmark processing time.
func BenchmarkWordProcessing(b *testing.B) {
	testChecker = reInit()
	for n := 0; n < b.N; n++ {
		testChecker.CheckFile(root, os.Getenv("GOPATH")+"/src/github.com/sudo-sturbia/gocheck/test/paragraph.txt")
	}
}
