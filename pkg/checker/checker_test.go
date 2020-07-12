// Package checker implements functions used to find spelling errors
// in a given text file and print error messages accordingly.
package checker

import (
	"os"
	"testing"

	"github.com/sudo-sturbia/gocheck/pkg/loader"
)

// Root Node of loaded dictionary.
var root *loader.Node

// Setup dictionary before testing.
func TestMain(m *testing.M) {
	root = loader.LoadDictionary("../../test/test-words.txt")
	os.Exit(m.Run())
}

// Test check file function on a file without errors.
func TestCheckFileWithoutErrors(t *testing.T) {
	testChecker := New()
	testChecker.CheckFile(root, "../../test/paragraph.txt")

	if len(testChecker.errors) != 0 {
		t.Errorf("Number of spelling errors incorrect, expected: 0, got: %d", len(testChecker.errors))
	}
}

// Test check file function on a file with errors.
func TestCheckFileWithErrors(t *testing.T) {
	testChecker := New()
	testChecker.CheckFile(root, "../../test/wrong-paragraph.txt")

	if len(testChecker.errors) != 8 {
		t.Errorf("Number of spelling errors incorrect, expected: 8, got: %d", len(testChecker.errors))
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
		if !errorMap[testChecker.errors[i]] {
			t.Errorf("Found incorrect error: %s", testChecker.errors[i])
		}
	}

}

// Benchmark processing time.
func BenchmarkWordProcessing(b *testing.B) {
	testChecker := New()
	for n := 0; n < b.N; n++ {
		testChecker.CheckFile(root, "../../test/paragraph.txt")
	}
}
