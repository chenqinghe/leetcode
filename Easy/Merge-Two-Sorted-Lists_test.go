package leetcode

import (
	"testing"
)

func TestMrgeTwoLists(t *testing.T) {

	var n2 *ListNode = &ListNode{3, nil}
	var n1 *ListNode = &ListNode{1, n2}

	var m2 *ListNode = &ListNode{4, nil}
	var m1 *ListNode = &ListNode{2, m2}

	if MrgeTwoLists(n1, m1) != n1 {
		t.Error("测试失败", MrgeTwoLists(n1, m1))
	}

}
