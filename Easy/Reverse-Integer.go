package leetcode

import (
	"math"
	"strconv"
)

func Reverse(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x = 0 - x //转换为正数
	}
	nstr := strconv.Itoa(x)
	nbyte := []byte(nstr)
	length := len(nbyte)
	after := make([]byte, length)
	for i := 0; i < length; i++ {
		after[length-1-i] = nbyte[i]
	}
	afternum, _ := strconv.Atoi(string(after))
	if afternum > math.MaxInt32 || afternum < math.MinInt32 {
		return 0
	}
	if neg {
		return 0 - afternum
	}
	return afternum
}
