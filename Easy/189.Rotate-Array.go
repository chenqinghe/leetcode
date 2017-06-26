package leetcode

import "fmt"

func Rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	fmt.Println(nums)
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
	fmt.Println(nums)
}

func reverse(nums []int, start int, end int) {
	for start < end {
		var tmp int
		tmp = nums[start]
		nums[start] = nums[end]
		nums[end] = tmp
		start++
		end--
	}
}
