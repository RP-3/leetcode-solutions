// https://leetcode.com/problems/sort-list/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	leftHead, rightHead := splitList(head)
	leftSorted, rightSorted := sortList(leftHead), sortList(rightHead)

	return zipLists(leftSorted, rightSorted)
}

func splitList(head *ListNode) (*ListNode, *ListNode) {
	if head == nil {
		return nil, nil
	}

	slow, fast := head, head.Next

	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	rightHead := slow.Next
	slow.Next = nil
	return head, rightHead
}

func zipLists(l, r *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for l != nil || r != nil {
		switch {
		case l == nil:
			tail.Next = r
			r = r.Next
		case r == nil:
			tail.Next = l
			l = l.Next
		case l.Val < r.Val:
			tail.Next = l
			l = l.Next
		default:
			tail.Next = r
			r = r.Next
		}
		tail = tail.Next
	}
	tail.Next = nil

	return dummy.Next
}
