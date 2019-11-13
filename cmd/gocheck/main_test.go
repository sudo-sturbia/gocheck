package main

import (
    "testing"
)

// Test check file function on several files
func TestCheckFile(t *testing.T) {
    root := loadDictionary("../../test/test_words.txt")

    // Test a file with no errors
    checkFile(root, "../../test/paragraph.txt")
    wg.Wait()
    if len(spellingErrors) != 0 {
        t.Errorf("Number of spelling errors incorrect, expected: 0, got: %d", len(spellingErrors))
    }

    // Test a file with eight errors
    checkFile(root, "../../test/paragraph-wrong.txt")
    wg.Wait()
    if len(spellingErrors) != 8 {
        t.Errorf("Number of spelling errors incorrect, expected: 8, got: %d", len(spellingErrors))
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
        if !errorMap[spellingErrors[i]] {
            t.Errorf("Found incorrect error: %s", spellingErrors[i])
        }
    }
}

// Benchmark word processing
func BenchmarkCheckFile(b *testing.B) {
    root := loadDictionary("../../test/test_words.txt")

    // Run benchmark
    for n := 0; n < b.N; n++ {
        checkFile(root, "../../test/paragraph.txt")
    }
}
