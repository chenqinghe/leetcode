package leetcode

import "testing"

func TestRemoveElement(t *testing.T) {
	if c := RemoveElement([]int{1, 2, 3, 4, 1}, 1); c != 3 {
		t.Fatal("testing failed.")
	}
	if c := RemoveElement([]int{3, 2, 2, 3, 4}, 3); c != 3 {
		t.Fatal("testing failed.")
	}
	t.Log("testing passed.")
}
