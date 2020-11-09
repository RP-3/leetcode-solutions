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

// abs returns the absolute value of its input
func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// ListNode is the standard leetcode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode is a standard leetcode definition for a binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
