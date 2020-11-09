// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxAncestorDiff(root *TreeNode) int {
	result := 0

	var minimax func(node *TreeNode) (int, int)
	minimax = func(node *TreeNode) (int, int) {
		lMin, lMax, rMin, rMax := node.Val, node.Val, node.Val, node.Val

		if node.Left != nil {
			lMin, lMax = minimax(node.Left)
		}
		if node.Right != nil {
			rMin, rMax = minimax(node.Right)
		}

		childMin, childMax := min(lMin, rMin), max(lMax, rMax)
		localResult := max(abs(childMin-node.Val), abs(childMax-node.Val))
		result = max(result, localResult)

		return min(node.Val, childMin), max(node.Val, childMax)
	}

	minimax(root)
	return result
}
