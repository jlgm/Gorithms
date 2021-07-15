//https://leetcode.com/problems/copy-list-with-random-pointer/
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	m, n, x := make(map[*Node]*Node), make(map[*Node]*Node), make(map[*Node]*Node)
	it := head
	for it != nil {
		m[it] = it.Random
		x[it] = it.Next
		it = it.Next
	}

	it = head
	for it != nil {
		n[it] = &Node{
			Val: it.Val,
		}
		it = it.Next
	}

	it = head
	for it != nil {
		n[it].Random = n[m[it]]
		n[it].Next = n[x[it]]
		it = it.Next
	}

	return n[head]
}
