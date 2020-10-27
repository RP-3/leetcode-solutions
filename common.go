package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ListNode is the standard leetcode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
