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

func (lis *LIS) Add(num int) int {
	idx := sort.SearchInts(lis.bucket, num)
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
