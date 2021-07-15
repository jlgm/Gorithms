//https://leetcode.com/problems/reverse-nodes-in-k-group/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func listSize(node *ListNode) int {
	it := node
	ans := 0
	for it != nil {
		it = it.Next
		ans += 1
	}
	return ans
}

func rev(node *ListNode, head **ListNode, tail **ListNode, k int) {
	if k == 1 {
		*head = node
		*tail = node.Next
		return
	}
	rev(node.Next, head, tail, k-1)
	node.Next.Next = node
	node.Next = *tail
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	size := listSize(head)
	curHead := head

	list := &ListNode{}
	listLast := list

	for i := 0; i < size/k; i++ {
		var h *ListNode
		var t *ListNode
		rev(curHead, &h, &t, k)

		it := h
		for j := 0; j < k; j++ {
			listLast.Next = &ListNode{
				Val: it.Val,
			}
			it = it.Next
			listLast = listLast.Next
		}
		curHead = t
	}

	for curHead != nil {
		listLast.Next = &ListNode{
			Val: curHead.Val,
		}
		curHead = curHead.Next
		listLast = listLast.Next
	}

	return list.Next
}