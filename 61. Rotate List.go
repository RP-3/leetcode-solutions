// https://leetcode.com/problems/rotate-list/

package leetcode

// ListNode is the standard leetcode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	length, tmp := 0, head
	for tmp != nil {
		length++
		tmp = tmp.Next
	}

	k %= length
	if k == 0 {
		return head
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	newHead := slow.Next
	slow.Next = nil
	fast.Next = head
	return newHead
}
