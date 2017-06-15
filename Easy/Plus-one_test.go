package leetcode

import (
	"fmt"
	"testing"
)

func TestPlusOne(t *testing.T) {
	is := PlusOne([]int{1, 2, 3})
	fmt.Println(is)
	is = PlusOne([]int{0, 1, 2, 3})
	fmt.Println(is)
	is = PlusOne([]int{0})
	fmt.Println(is)
	is = PlusOne([]int{9, 9})
	fmt.Println(is)
	t.Log("testing passed.")
}
