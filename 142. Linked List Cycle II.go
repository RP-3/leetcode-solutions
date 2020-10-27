// https://leetcode.com/problems/linked-list-cycle-ii/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	nodes := make(map[*ListNode]bool)
	p := head

	for p != nil {
		_, found := nodes[p]
		if found {
			return p
		}
		nodes[p] = true
		p = p.Next
	}

	return p
}
