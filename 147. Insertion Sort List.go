// https://leetcode.com/problems/insertion-sort-list/

package leetcode

import "math"

func insertionSortList(head *ListNode) *ListNode {
	result := &ListNode{Val: math.MaxInt32}

	curr := head
	for curr != nil {
		tmp, p := curr.Next, result

		for p.Next != nil && curr.Val > p.Next.Val {
			p = p.Next
		}

		curr.Next = p.Next
		p.Next = curr

		curr = tmp
	}

	return result.Next
}
