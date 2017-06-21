package leetcode

import "testing"

func TestMaxSubArray1(t *testing.T) {
	if i := MaxSubArray([]int{1, 2, 3}); i != 6 {
		t.Fatal("testing failed.", i)
	}
	t.Log("testing passed.")
}

func TestMaxSubArray2(t *testing.T) {
	if i := MaxSubArray([]int{-1, 2, 3}); i != 5 {
		t.Fatal("testing failed.", i)
	}
	t.Log("testing passed.")
}

func TestMaxSubArray3(t *testing.T) {
	if i := MaxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}); i != 6 {
		t.Fatal("testing failed.", i)
	}
	t.Log("testing passed.")
}

func TestMaxSubArray4(t *testing.T) {
	if i := MaxSubArray([]int{0}); i != 0 {
		t.Fatal("testing failed.", i)
	}
	t.Log("testing passed.")
}

func TestMaxSubArray5(t *testing.T) {
	if i := MaxSubArray([]int{-1, 1}); i != 1 {
		t.Fatal("testing failed.", i)
	}
	t.Log("testing passed.")
}

func TestMaxSubArray6(t *testing.T) {
	if i := MaxSubArray([]int{1, -1, 2}); i != 2   {
		t.Fatal("testing failed.", i)
	}
	t.Log("testing passed.")
}
