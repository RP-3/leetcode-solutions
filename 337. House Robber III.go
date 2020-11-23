// https://leetcode.com/problems/house-robber-iii/

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rob(root *TreeNode) int {
	// returns a tuple containing the max return from robbing
	// [0] this house and [1] not robbing this house
	var inner func(*TreeNode) (int, int)
	inner = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		lRob, lNoRob := inner(node.Left)
		rRob, rNoRob := inner(node.Right)
		maxRobbingThis := node.Val + lNoRob + rNoRob
		maxNotRobbingThis := max(lRob, lNoRob) + max(rRob, rNoRob)
		return maxRobbingThis, maxNotRobbingThis
	}

	robRoot, noRobRoot := inner(root)
	return max(robRoot, noRobRoot)
}
