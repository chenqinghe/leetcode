package leetcode

import "fmt"

func SearchInsert(nums []int, target int) int {
	l := len(nums)
	if l <= 1 { //只有一个元素情况
		if target <= nums[0] {
			return 0
		}
		return 1
	}
	//多个元素情况
	index := 0
	for i := 1; i <= l-1; i++ { //从1开始，l-1截止防止越界
		fmt.Println(i, nums[i-1], nums[i])
		if nums[i-1] >= target {
			index = i - 1
			break
		}
		if target > nums[i-1] && target <= nums[i] {
			index = i
			break
		}
		if target > nums[i] && i >= (l-1) { //
			index = l
		}
	}
	return index
}
