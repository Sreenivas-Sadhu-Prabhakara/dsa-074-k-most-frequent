package main

import (
	"reflect"
	"testing"
)

func TestKMostFrequent(t *testing.T) {
	cases := []struct {
		in   []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{4, 4, 7, 7}, 2, []int{4, 7}}, // tie broken toward smaller value
		{[]int{5, 5, 5}, 1, []int{5}},
		{[]int{1, 2}, 5, []int{1, 2}}, // k larger than distinct count
		{[]int{}, 2, []int{}},
	}
	for _, c := range cases {
		if got := KMostFrequent(c.in, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("KMostFrequent(%v,%d) = %v, want %v", c.in, c.k, got, c.want)
		}
	}
}
