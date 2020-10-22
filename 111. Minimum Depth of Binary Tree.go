// https://leetcode.com/problems/minimum-depth-of-binary-tree/

package leetcode

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := math.MaxInt32

	var min func(*TreeNode, int)
	min = func(node *TreeNode, depth int) {
		switch {
		case node == nil: // no root
			return
		case node.Left == nil && node.Right == nil: // found a leaf
			if depth < result {
				result = depth
			}
		case depth > result: // no point looking further
			return
		default: // keep looking
			min(node.Left, depth+1)
			min(node.Right, depth+1)
		}
	}

	min(root, 1)
	return result
}
