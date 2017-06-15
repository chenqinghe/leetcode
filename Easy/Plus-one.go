package leetcode

func PlusOne(digits []int) []int {
	l := len(digits)
	if digits[l-1] < 9 {
		digits[l-1]++
		return digits
	}
	if l-1 >= 0 {
		digits[l-1] = 0
		digits[l-2]++
		return digits
	} else {
		digits[l-1] = 0
		digits = append([]int{1}, digits...)
	}
	return digits
}
