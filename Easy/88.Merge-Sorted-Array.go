package leetcode

import (
	"fmt"
	"sort"
)

//That means, m+n is no bigger that the size of nums1.
// And m is not the size of nums1 but the number of elements used in nums1,
// and so is n.

//---------------------way 1---------------------
func MergeSortedArray(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[:m]
	nums2 = nums2[:n]
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
}

//-------------------------way 2---------------------------
