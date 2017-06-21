package leetcode

import "fmt"

//思路：从头尾向中部“夹逼”
func MaxSubArray(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return nums[0]
	}
	head := 0
	foot := l
	headsubarrsum := 0
	footsubarrsum := 0
	for i := 0; i < l; i++ {
		if i >= l-i { //遍历完成
			break
		}
		headsubarrsum += nums[i]
		footsubarrsum += nums[l-1-i]
		if headsubarrsum <= 0 { //子串和小于零，抛弃
			head = i + 1
			headsubarrsum = 0
		}
		if footsubarrsum <= 0 { //子串和小于零，抛弃
			foot = l - i - 1
			footsubarrsum = 0
		}
		fmt.Println(head, foot)
	}

	nums = nums[head:foot]
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
