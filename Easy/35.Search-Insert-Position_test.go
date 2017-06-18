package leetcode

import "testing"

func TestSearchInsert1(t *testing.T) {
	if i := SearchInsert([]int{1, 3, 5, 6}, 5); i != 2 {
		t.Fatal("testing1 failed.", i)
	}
	t.Log("testing passed.")

}

func TestSearchInsert2(t *testing.T) {

	if i := SearchInsert([]int{1, 3, 5, 6}, 2); i != 1 {
		t.Fatal("testing2 failed.", i)
	}
	t.Log("testing passed.")

}
func TestSearchInsert3(t *testing.T) {

	if i := SearchInsert([]int{1, 3, 5, 6}, 7); i != 4 {
		t.Fatal("testing3 failed.", i)
	}
	t.Log("testing passed.")

}
func TestSearchInsert4(t *testing.T) {

	if i := SearchInsert([]int{1, 3, 5, 6}, 0); i != 0 {
		t.Fatal("testing4 failed.", i)
	}
	t.Log("testing passed.")

}
func TestSearchInsert5(t *testing.T) {
	if i := SearchInsert([]int{1}, 1); i != 0 {
		t.Fatal("testing5 failed.", i)
	}
	t.Log("testing passed.")

}
