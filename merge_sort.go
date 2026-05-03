package sort

import (
	"math/bits"
	"runtime"
	"sync"
)

func mergeSl(sl1 []int, sl2 []int) []int {
	r := make([]int, 0, len(sl1)+len(sl2))
	i1 := 0
	i2 := 0
	for i1 < len(sl1) && i2 < len(sl2) {
		if sl1[i1] <= sl2[i2] {
			r = append(r, sl1[i1])
			i1++
		} else {
			r = append(r, sl2[i2])
			i2++
		}
	}
	for ; i1 < len(sl1); i1++ {
		r = append(r, sl1[i1])
	}
	for ; i2 < len(sl2); i2++ {
		r = append(r, sl2[i2])
	}
	return r
}

func InsertionSort(sl []int) {
	for i := 1; i < len(sl); i++ {
		for j := i - 1; j >= 0; j-- {
			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			} else {
				break
			}
		}
	}
}

const insertThreshold = 32

func MergeSort(sl []int) []int {
	if len(sl) <= insertThreshold {
		InsertionSort(sl)
		return sl
	}
	sl1 := MergeSort(sl[:len(sl)/2])
	sl2 := MergeSort(sl[len(sl)/2:])
	r := mergeSl(sl1, sl2)
	return r
}

const concurThreshold = 2048

func MergeSortConcurrent(sl []int) []int {
	if len(sl) < concurThreshold {
		return MergeSort(sl)
	}
	wg := sync.WaitGroup{}
	var sl1, sl2 []int
	wg.Go(func() {
		sl1 = MergeSortConcurrent(sl[:len(sl)/2])
	})
	sl2 = MergeSortConcurrent(sl[len(sl)/2:])
	wg.Wait()
	return mergeSl(sl1, sl2)
}

func MergeSortSemaphore(sl []int) []int {
	sem := make(chan struct{}, runtime.NumCPU()*4)
	return mergeSortSemaphore(sem, sl)
}
func mergeSortSemaphore(sem chan struct{}, sl []int) []int {
	if len(sl) < concurThreshold {
		return MergeSort(sl)
	}
	wg := sync.WaitGroup{}
	var sl1, sl2 []int
	select {
	case sem <- struct{}{}:
		wg.Go(func() {
			sl1 = mergeSortSemaphore(sem, sl[:len(sl)/2])
			<-sem
		})
	default:
		sl1 = mergeSortSemaphore(sem, sl[:len(sl)/2])
	}
	sl2 = mergeSortSemaphore(sem, sl[len(sl)/2:])
	wg.Wait()
	return mergeSl(sl1, sl2)
}

func MergeSortConcurrentDepth(sl []int) []int {
	depth := bits.Len(uint(runtime.NumCPU())) - 1
	return mergeSortConcurrentDepth(sl, depth)
}

func mergeSortConcurrentDepth(sl []int, depth int) []int {
	if len(sl) < concurThreshold || depth == 0 {
		return MergeSort(sl)
	}
	wg := sync.WaitGroup{}
	var sl1, sl2 []int
	wg.Go(func() {
		sl1 = mergeSortConcurrentDepth(sl[:len(sl)/2], depth-1)
	})
	sl2 = mergeSortConcurrentDepth(sl[len(sl)/2:], depth-1)
	wg.Wait()
	return mergeSl(sl1, sl2)

}
