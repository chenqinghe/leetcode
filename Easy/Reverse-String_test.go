package leetcode

import "testing"

func TestReverseString(t *testing.T) {
	if s := ReverseStringOne("hello"); s != "olleh" {
		t.Fatal("testing failed.")
	}
	if s := ReverseStringOne("  he"); s != "eh  " {
		t.Fatal("testing failed: string that contains space reverse incorrectly.")
	}
	t.Log("testing passed.")
}

func TestReverseStringTwo(t *testing.T) {
	if s := ReverseStringTwo("hello"); s != "olleh" {
		t.Fatal("testing failed.")
	}
	if s := ReverseStringTwo("  he"); s != "eh  " {
		t.Fatal("testing failed: string that contains space reverse incorrectly.")
	}
	t.Log("testing passed.")
}

func TestReverseStringThree(t *testing.T) {
	if s := ReverseStringThree("hello"); s != "olleh" {
		t.Fatal("testing failed.")
	}
	if s := ReverseStringThree("  he"); s != "eh  " {
		t.Fatal("testing failed: string that contains space reverse incorrectly.")
	}
	t.Log("testing passed.")
}
