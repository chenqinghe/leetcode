package leetcode

import "fmt"

func PlusOne(digits []int) []int {
	l := len(digits)
	//不进位情况
	if digits[l-1] < 9 {
		fmt.Print(l - 1)
		digits[l-1]++
		return digits
	}
	//有进位情况
	if l-2 >= 0 { //前一位存在
		fmt.Println(l - 1)
		digits = append(PlusOne(digits[:l-1]), 0)
	} else { //前一位不存在
		digits = []int{1, 0}
	}
	return digits
}
