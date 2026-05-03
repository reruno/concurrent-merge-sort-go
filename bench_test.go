package sort_test

import (
	"fmt"
	"math/rand/v2"
	"testing"

	sort "github.com/reruno/concurrent-merge-sort-go"
)

var insertionSizes = []int{100, 1_000, 10_000}
var mergeSortSizes = []int{100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000}

func makeRandom(n int) []int {
	sl := make([]int, n)
	for i := range sl {
		sl[i] = rand.IntN(1_000_000)
	}
	return sl
}

func formatSize(n int) string {
	suffixes := []struct {
		threshold int
		symbol    string
	}{
		{1_000_000_000, "B"},
		{1_000_000, "M"},
		{1_000, "K"},
	}
	for _, s := range suffixes {
		if n%s.threshold == 0 {
			return fmt.Sprintf("%d%s", n/s.threshold, s.symbol)
		}
	}
	return fmt.Sprintf("%d", n)
}

func BenchmarkInsertionSort(b *testing.B) {
	for _, n := range insertionSizes {
		b.Run(formatSize(n), func(b *testing.B) {
			src := makeRandom(n)
			b.ResetTimer()
			for b.Loop() {
				sl := make([]int, n)
				copy(sl, src)
				sort.InsertionSort(sl)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for _, n := range mergeSortSizes {
		b.Run(formatSize(n), func(b *testing.B) {
			src := makeRandom(n)
			b.ResetTimer()
			for b.Loop() {
				sl := make([]int, n)
				copy(sl, src)
				sort.MergeSort(sl)
			}
		})
	}
}

func BenchmarkMergeSortConcurrent(b *testing.B) {
	for _, n := range mergeSortSizes {
		b.Run(formatSize(n), func(b *testing.B) {
			src := makeRandom(n)
			b.ResetTimer()
			for b.Loop() {
				sl := make([]int, n)
				copy(sl, src)
				sort.MergeSortConcurrent(sl)
			}
		})
	}
}

func BenchmarkMergeSortConcurrentDepth(b *testing.B) {
	for _, n := range mergeSortSizes {
		b.Run(formatSize(n), func(b *testing.B) {
			src := makeRandom(n)
			b.ResetTimer()
			for b.Loop() {
				sl := make([]int, n)
				copy(sl, src)
				sort.MergeSortConcurrentDepth(sl)
			}
		})
	}
}

func BenchmarkMergeSortSemaphore(b *testing.B) {
	for _, n := range mergeSortSizes {
		b.Run(formatSize(n), func(b *testing.B) {
			src := makeRandom(n)
			b.ResetTimer()
			for b.Loop() {
				sl := make([]int, n)
				copy(sl, src)
				sort.MergeSortSemaphore(sl)
			}
		})
	}
}
