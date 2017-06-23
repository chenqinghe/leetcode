package leetcode

import (
	"sort"
)

func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		nums1 = nums2
		return
	}
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
}
