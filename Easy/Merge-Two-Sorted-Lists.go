package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//lists 按从小到大排序
//思路：遍历l2，将l2合并到l1中
func MrgeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	for {
		if l2.Next != nil {
			t1 := l2
			l2 = l2.Next
			t1.Next = nil
			return MrgeTwoLists(MrgeTwoLists(l1, t1), l2)
		} else { //L2尾节点
			tmp := l1
			for {
				if tmp.Next == nil { //L1 单节点
					if tmp.Val > l2.Val {
						l2.Next = tmp
						l1 = l2
					} else {
						tmp.Next = l2
					}
					break //处理完毕，退出
				} else {
					if l2.Val > tmp.Val {
						if l2.Val > tmp.Next.Val {
							tmp = tmp.Next
							continue
						} else {
							l2.Next = tmp.Next
							tmp.Next = l2
						}
					} else {
						l2.Next = tmp.Next
						tmp.Next = l2
					}
				}
			}
		}
	}
	return l1
}
