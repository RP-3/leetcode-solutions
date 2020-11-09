// https://leetcode.com/problems/binary-tree-tilt/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	var sumAndTilt func(node *TreeNode) (int, int)
	sumAndTilt = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}

		lSum, lTilt := sumAndTilt(node.Left)
		rSum, rTilt := sumAndTilt(node.Right)
		sum := lSum + rSum + node.Val
		tilt := abs(rSum - lSum)

		return sum, lTilt + rTilt + tilt
	}

	_, tilt := sumAndTilt(root)
	return tilt
}
