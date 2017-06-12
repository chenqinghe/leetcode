package leetcode

//时间复杂度:O(n)
//空间复杂度：O(1)
func LengthOfLastWord(s string) int {
	l := len(s)
	length := 0
	flag := false
	for i := 0; i < l; i++ {
		if s[i] != ' ' {
			if flag {
				length = 1
				flag = false
			} else {
				length++
			}
			continue
		}
		flag = true
	}
	return length
}
