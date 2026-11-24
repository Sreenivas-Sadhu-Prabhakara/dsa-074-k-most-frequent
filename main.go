package main

import (
	"fmt"
	"sort"
)

// KMostFrequent returns the k values with the highest occurrence counts.
// Ties break toward the smaller value so the result is deterministic.
// Bucket sort on frequency: counts are bounded by n, so no heap is needed.
func KMostFrequent(nums []int, k int) []int {
	if k < 1 {
		return []int{}
	}
	count := map[int]int{}
	for _, v := range nums {
		count[v]++
	}
	// buckets[f] = values occurring exactly f times
	buckets := make([][]int, len(nums)+1)
	for v, f := range count {
		buckets[f] = append(buckets[f], v)
	}
	out := []int{}
	for f := len(buckets) - 1; f >= 1 && len(out) < k; f-- {
		sort.Ints(buckets[f])
		for _, v := range buckets[f] {
			out = append(out, v)
			if len(out) == k {
				break
			}
		}
	}
	return out
}

func main() {
	fmt.Println(KMostFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
