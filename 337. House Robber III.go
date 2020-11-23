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
	memo := make(map[robberCacheKey]int)
	var inner func(*TreeNode, bool) int

	inner = func(node *TreeNode, canRob bool) int {
		key := robberCacheKey{node: node, canRob: canRob}
		cached, ok := memo[key]
		switch {
		case node == nil:
			return 0
		case ok:
			return cached
		case canRob:
			memo[key] = max(
				node.Val+inner(node.Left, false)+inner(node.Right, false), // rob
				inner(node.Left, true)+inner(node.Right, true),            // do not rob
			)
		default: // cannot rob
			memo[key] = inner(node.Left, true) + inner(node.Right, true) // do not rob
		}
		return memo[key]
	}

	return max(inner(root, false), inner(root, true))
}

type robberCacheKey struct {
	node   *TreeNode
	canRob bool
}
