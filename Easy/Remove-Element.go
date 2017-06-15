package leetcode

//时间复杂度：O(n)
//空间复杂度: O(1)
func RemoveElement(nums []int, val int) int {
	count := len(nums)
	index := 0 //计数变量
	for i := 0; i < count; i++ {
		if nums[i] != val { //不等时，放入当前数组，相等时忽略
			nums[index] = nums[i]
			index++
		}
	}
	return index
}
