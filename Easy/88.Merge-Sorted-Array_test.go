package leetcode

import "testing"

func TestMergeSortedArray(t *testing.T) {
	MergeSortedArray([]int{1, 2, 3}, 3, []int{4, 5, 6}, 3)
}

func TestMergeSortedArray2(t *testing.T) {
	MergeSortedArray([]int{0}, 0, []int{1}, 1)
}
