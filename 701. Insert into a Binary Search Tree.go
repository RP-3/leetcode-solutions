// https://leetcode.com/problems/insert-into-a-binary-search-tree/solution/

package leetcode

// TreeNode is a standard leetcode binary tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	switch {
	case root == nil:
		return &TreeNode{val, nil, nil}
	case val < root.Val:
		root.Left = insertIntoBST(root.Left, val)
	default:
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}

// Bonus, itreative version. No call stack, but also no faster ¯\_(ツ)_/¯
func iterativeInsertIntoBST(root *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{val, nil, nil}
	if root == nil {
		return newNode
	}

	curr, left := root, false
	var prev *TreeNode = nil
	for curr != nil {
		prev = curr
		if val < curr.Val {
			curr = curr.Left
			left = true
		} else {
			curr = curr.Right
			left = false
		}
	}

	if left {
		prev.Left = newNode
	} else {
		prev.Right = newNode
	}

	return root
}
