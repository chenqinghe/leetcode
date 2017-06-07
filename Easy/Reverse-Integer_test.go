package leetcode

import (
	"testing"
)

func TestReverse(t *testing.T) {
	//普通测试
	if i := Reverse(123); i != 321 {
		t.Error("测试失败")
	}
	//测试负数
	if i := Reverse(-567); i != -765 {
		t.Error("负数测试失败", i)
	}
	//测试是否超过限制
	if i := Reverse(1000000003); i != 0 {
		t.Error("function returns 0 when the reversed integer overflows.")
	}
	//测试先导0
	if i := Reverse(100); i != 1 {
		t.Error("prev zero should be ignored")
	}
	t.Log("test passed.")
}
