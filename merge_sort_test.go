package sort_test

import (
	"math/rand/v2"
	"slices"
	"testing"

	sort "github.com/reruno/concurrent-merge-sort-go"
)

func TestMergeSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"two sorted", []int{1, 2}, []int{1, 2}},
		{"two reversed", []int{2, 1}, []int{1, 2}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},
		{"all same", []int{7, 7, 7}, []int{7, 7, 7}},
		{"negatives", []int{-3, 1, -1, 2, 0}, []int{-3, -1, 0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sort.MergeSort(tt.input)
			if !slices.Equal(got, tt.want) {
				t.Errorf("MergeSort(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMergeSortLarge(t *testing.T) {
	input := rand.N(10_000)
	sl := make([]int, input)
	for i := range sl {
		sl[i] = rand.IntN(1_000_000)
	}
	want := slices.Clone(sl)
	slices.Sort(want)

	got := sort.MergeSort(sl)
	if !slices.Equal(got, want) {
		t.Errorf("MergeSort failed on large random input of size %d", input)
	}
}

func TestMergeSortSemaphore(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"two sorted", []int{1, 2}, []int{1, 2}},
		{"two reversed", []int{2, 1}, []int{1, 2}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},
		{"all same", []int{7, 7, 7}, []int{7, 7, 7}},
		{"negatives", []int{-3, 1, -1, 2, 0}, []int{-3, -1, 0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sort.MergeSortSemaphore(tt.input)
			if !slices.Equal(got, tt.want) {
				t.Errorf("MergeSortSemaphore(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMergeSortSemaphoreLarge(t *testing.T) {
	input := rand.N(10_000)
	sl := make([]int, input)
	for i := range sl {
		sl[i] = rand.IntN(1_000_000)
	}
	want := slices.Clone(sl)
	slices.Sort(want)

	got := sort.MergeSortSemaphore(sl)
	if !slices.Equal(got, want) {
		t.Errorf("MergeSortSemaphore failed on large random input of size %d", input)
	}
}

func TestMergeSortConcurrentDepth(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"two sorted", []int{1, 2}, []int{1, 2}},
		{"two reversed", []int{2, 1}, []int{1, 2}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},
		{"all same", []int{7, 7, 7}, []int{7, 7, 7}},
		{"negatives", []int{-3, 1, -1, 2, 0}, []int{-3, -1, 0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sort.MergeSortConcurrentDepth(tt.input)
			if !slices.Equal(got, tt.want) {
				t.Errorf("MergeSortConcurrentDepth(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestMergeSortConcurrentDepthLarge(t *testing.T) {
	input := rand.N(10_000)
	sl := make([]int, input)
	for i := range sl {
		sl[i] = rand.IntN(1_000_000)
	}
	want := slices.Clone(sl)
	slices.Sort(want)

	got := sort.MergeSortConcurrentDepth(sl)
	if !slices.Equal(got, want) {
		t.Errorf("MergeSortConcurrentDepth failed on large random input of size %d", input)
	}
}

func TestInsertionSort(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"two sorted", []int{1, 2}, []int{1, 2}},
		{"two reversed", []int{2, 1}, []int{1, 2}},
		{"already sorted", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"reversed", []int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{"duplicates", []int{3, 1, 2, 1, 3}, []int{1, 1, 2, 3, 3}},
		{"all same", []int{7, 7, 7}, []int{7, 7, 7}},
		{"negatives", []int{-3, 1, -1, 2, 0}, []int{-3, -1, 0, 1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := make([]int, len(tt.input))
			copy(got, tt.input)
			sort.InsertionSort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("InsertionSort(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
