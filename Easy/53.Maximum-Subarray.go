package leetcode


func MaxSubArray(nums []int) int {

	//-------------way 1-------------------
	////思路：从头尾向中部“夹逼”

	//l := len(nums)
	//if l <= 1 {
	//	return nums[0]
	//}
	//head := 0
	//foot := l
	//headsubarrsum := 0
	//footsubarrsum := 0
	//for i := 0; i < l; i++ {
	//	if i >= l-i { //遍历完成
	//		break
	//	}
	//	headsubarrsum += nums[i]
	//	footsubarrsum += nums[l-1-i]
	//	if headsubarrsum <= 0 { //子串和小于零，抛弃
	//		head = i + 1
	//		headsubarrsum = 0
	//	}
	//	if footsubarrsum <= 0 { //子串和小于零，抛弃
	//		foot = l - i - 1
	//		footsubarrsum = 0
	//	}
	//	fmt.Println(head, foot)
	//}
	//
	//nums = nums[head:foot]
	//sum := 0
	//for _, v := range nums {
	//	sum += v
	//}
	//return sum

	//---------------way 2-----------------------
	maxsofar := nums[0]
	maxEndingHere := nums[0]
	for i := 1; i < len(nums); i++ {
		maxEndingHere = max(maxEndingHere+nums[i], nums[i])
		maxsofar = max(maxEndingHere, maxsofar)
	}
	return maxsofar
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
