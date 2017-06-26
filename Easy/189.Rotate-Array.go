package leetcode

import "fmt"

func Rotate(nums []int, k int) {
	l := len(nums)
	k = k % l
	fmt.Println(nums)
	for i := 0; i < l; i++ {
		tmp := nums[(i+k)%l]
		nums[(i+k)%l] = nums[i]
	}
	fmt.Println(nums)
}
