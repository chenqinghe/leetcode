package leetcode

func RemoveDuplicates(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return 1
	}
	index := 1
	tmp := nums[0]
	for i := 1; i < l; i++ {
		if nums[i] != tmp {
			nums[index] = nums[i]
			tmp = nums[index]
			index++
		}
	}
	nums = nums[:index]
	return len(nums)
}
