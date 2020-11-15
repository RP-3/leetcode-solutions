// https://leetcode.com/problems/range-sum-of-bst/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, low int, high int) int {
	switch {
	case root == nil:
		return 0
	case root.Val < low:
		return rangeSumBST(root.Right, low, high)
	case root.Val > high:
		return rangeSumBST(root.Left, low, high)
	default:
		lSum, rSum := rangeSumBST(root.Left, low, high), rangeSumBST(root.Right, low, high)
		return lSum + rSum + root.Val
	}
}
