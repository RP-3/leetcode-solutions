// https://leetcode.com/problems/serialize-and-deserialize-bst/

package leetcode

import "math"

// Codec : Really not sure what I'm meant to do with this...
type Codec struct{}

// func Constructor() Codec {
// 	return Codec{}
// }

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	result := []rune{}

	var preorderTraversal func(node *TreeNode)
	preorderTraversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, rune(node.Val))
		preorderTraversal(node.Left)
		preorderTraversal(node.Right)
	}

	preorderTraversal(root)
	return string(result)
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	nodeData, i := []rune(data), 0

	var preorderTraversal func(min, max int) *TreeNode
	preorderTraversal = func(min, max int) *TreeNode {
		if i >= len(nodeData) {
			return nil
		}

		curr := int(nodeData[i])
		if curr > max || curr < min {
			return nil
		}

		i++
		newNode := &TreeNode{curr, nil, nil}
		newNode.Left = preorderTraversal(min, curr)
		newNode.Right = preorderTraversal(curr, max)
		return newNode
	}

	return preorderTraversal(math.MinInt32, math.MaxInt32)
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
