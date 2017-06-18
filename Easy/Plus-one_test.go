package leetcode

import (
	"testing"
)

func TestPlusOne(t *testing.T) {
	if is := PlusOne([]int{1, 2, 3}); !compareSlices_BCE(is, []int{1, 2, 4}) {
		t.Fatal("simple testing failed.result should be [1 2 4], but", is, "is given")
	}
	if is := PlusOne([]int{0, 1, 2, 3}); !compareSlices_BCE(is, []int{0, 1, 2, 4}) {
		t.Fatal("leading 0 testing failed.result should be [0 1 2 4], but", is, "is given")
	}
	if is := PlusOne([]int{0}); !compareSlices_BCE(is, []int{1}) {
		t.Fatal("0 testing failed.result should be [1], but", is, "is given")
	}
	if is := PlusOne([]int{9}); !compareSlices_BCE(is, []int{1, 0}) {
		t.Fatal("carry testing failed.result should be [1 0], but", is, "is given")
	}
	if is := PlusOne([]int{9, 9}); !compareSlices_BCE(is, []int{1, 0, 0}) {
		t.Fatal("carry testing failed.result should be [1 0 0], but", is, "is given")
	}
	t.Log("testing passed.")
}

func compareSlices_BCE(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	b = b[:len(a)] // this line is the key
	for i, v := range a {
		if v != b[i] { // here is no bounds checking for b[i]
			return false
		}
	}
	return true
}
