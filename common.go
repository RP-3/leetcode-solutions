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

// sort intervals by their start time
type byStartTime [][]int

func (a byStartTime) Len() int           { return len(a) }
func (a byStartTime) Less(i, j int) bool { return a[i][0] < a[j][0] }
func (a byStartTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
