package leetcode

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	if i := LengthOfLastWord("  "); i != 0 {
		t.Fatal("testing failed: length of last word should be 0")
	}
	if i := LengthOfLastWord("h l "); i != 1 {
		t.Fatal("testing failed: length of last word should be 1", i)
	}
	if i := LengthOfLastWord(" l"); i != 1 {
		t.Fatal("testing failed.")
	}
	if i := LengthOfLastWord("he ll os"); i != 2 {
		t.Fatal("testing failed.")
	}
	if i := LengthOfLastWord("a "); i != 1 {
		t.Fatal("testing failed.")
	}
	t.Log("testing passed.")
}
