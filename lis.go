package main

import (
	"sort"
)

// FIXME: beyond hardcoded capacity this will copy the slice over.
const BUCKET_CAPACITY = 500000

type LIS struct {
	ch     chan int
	bucket []int
}

func NewLIS() *LIS {
	return &LIS{make(chan int, 100), make([]int, 0, BUCKET_CAPACITY)}
}

func (lis *LIS) findIdx(num int) int {
	// on a mostly sorted array, check the last bucket before doing binary search.
	bLen := len(lis.bucket)
	if bLen > 0 && lis.bucket[bLen-1] < num {
		return bLen
	} else {
		return sort.SearchInts(lis.bucket, num)
	}
}

func (lis *LIS) Add(num int) int {
	idx := lis.findIdx(num)
	switch {
	case idx == len(lis.bucket):
		lis.bucket = append(lis.bucket, num)
	case lis.bucket[idx] == num:
		// Duplicate; ignore.
	default:
		lis.bucket[idx] = num
	}
	return len(lis.bucket)
}

func (lis *LIS) Len() int {
	return len(lis.bucket)
}
