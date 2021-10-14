/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(cur *ListNode, k int) (*ListNode, *ListNode) {
	if cur.Next == nil || k == 1 {
		return cur, cur
	}
	tmp, head := reverseList(cur.Next, k-1)
	tmp.Next = cur
	cur.Next = nil
	return cur, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {

	it, kk := head, k
	for it != nil && kk > 0 {
		kk--
		it = it.Next
	}
	if kk > 0 {
		return head //base case
	}

	tail, newHead := reverseList(head, k)
	tail.Next = reverseKGroup(it, k)

	return newHead
}