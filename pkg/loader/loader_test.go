package loader

import (
	"testing"
)

// Benchmark dictionary loading
func BenchmarkDictionaryLoading(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LoadDictionary("../../test/test_words.txt")
	}
}
