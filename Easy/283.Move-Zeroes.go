package leetcode

func MoveZeroes(nums []int) {
	l := len(nums)
	i := 0
	leftZero := -1 //最左边的0的位置
	for i < l-1 {
		if nums[i] != 0 {
			continue
		} else {
			if leftZero < 0 { //初始化
				leftZero = i
			}
			//交换位置
			nums[i], nums[leftZero] = nums[leftZero], nums[i]
			if countZero > 1 {
				leftZero++
			} else {
				leftZero
			}
		}
	}
}
