package leetcode

import "testing"

func TestRemoveDuplicates1(t *testing.T) {
	if i := RemoveDuplicates([]int{1}); i != 1 {
		t.Fatal("testing1 failed.", i)
	}
	t.Log("testing1 passed.")
}
func TestRemoveDuplicates2(t *testing.T) {
	if i := RemoveDuplicates([]int{1, 1, 3, 4, 5}); i != 4 {
		t.Fatal("testing2 failed.", i)
	}
	t.Log("testing2 passed.")

}
func TestRemoveDuplicates3(t *testing.T) {
	if i := RemoveDuplicates([]int{1, 1, 1, 1, 1}); i != 1 {
		t.Fatal("testing3 failed.", i)
	}
	t.Log("testing3 passed.")

}
func TestRemoveDuplicates4(t *testing.T) {
	if i := RemoveDuplicates([]int{1, 2, 3, 4, 5}); i != 5 {
		t.Fatal("testing4 failed.", i)
	}
	t.Log("testing4 passed.")

}

func TestRemoveDuplicates5(t *testing.T) {
	if i := RemoveDuplicates([]int{1, 1, 2, 2}); i != 2 {
		t.Fatal("testing5 failed.", i)
	}
	t.Log("testing5 passed.")

}

func TestRemoveDuplicates6(t *testing.T) {
	if i := RemoveDuplicates([]int{1, 2}); i != 2 {
		t.Fatal("testing6 failed.", i)
	}
	t.Log("testing6 passed.")

}
