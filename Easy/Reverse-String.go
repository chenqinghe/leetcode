package leetcode

import "bytes"

//方法1：使用字节数组
//时间复杂度：O(n)
//空间复杂度：O(n)
func ReverseStringOne(s string) string {
	tmp := []byte(s)
	l := len(tmp)
	var count int
	if l%2 == 0 {
		count = l / 2
	} else {
		count = (l - 1) / 2
	}
	for i := 0; i < count; i++ {
		t := tmp[i]
		tmp[i] = tmp[l-1-i]
		tmp[l-1-i] = t
	}
	return string(tmp)
}

//方法2：倒序遍历
//时间复杂度：O(n)
//空间复杂度：O(n)
func ReverseStringTwo(s string) string {
	var buf bytes.Buffer
	l := len(s)
	for i := 0; i < l; i++ {
		buf.WriteByte(s[l-1-i])
	}
	return buf.String()
}

//方法3：异或位运算
//时间复杂度：O(n)
//空间复杂度：O(1)
//使用异或交换变量值：
//A=A^B
//B=A^B
//A=A^B
//通过以上三步，能实现在程序中交换两个变量的数值的目标。
func ReverseStringThree(s string) string {
	tmp := []byte(s)
	l := len(tmp)
	var count int
	if l%2 == 0 {
		count = l / 2
	} else {
		count = (l - 1) / 2
	}
	for i := 0; i < count; i++ {
		tmp[i] = tmp[i] ^ tmp[l-1-i]
		tmp[l-1-i] = tmp[i] ^ tmp[l-1-i]
		tmp[i] = tmp[i] ^ tmp[l-1-i]
	}
	return string(tmp)
}
