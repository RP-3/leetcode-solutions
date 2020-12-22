package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minIntSlice(s []int) int {
	result := s[0]
	for _, n := range s {
		result = min(result, n)
	}
	return result
}

func sumIntSlice(s []int) int {
	result := 0
	for _, n := range s {
		result += n
	}
	return result
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

// pow returns a ^ b, where b is a non-negative integer
func pow(a, b int) int {
	switch {
	case b < 0:
		panic("We're just handling positive exponents here")
	case b == 0:
		return 1
	case b == 1:
		return a
	case b%2 == 0:
		return pow(a*a, b/2)
	default:
		return a * pow(a*a, (b-1)/2)
	}
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
