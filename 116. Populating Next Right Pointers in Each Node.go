// https://leetcode.com/problems/populating-next-right-pointers-in-each-node/

package leetcode

// TreeNodeWithNext is a non-standard binary tree node
// that also has a pointer "next" to its right sibling/cousin
type TreeNodeWithNext struct {
	Val   int
	Left  *TreeNodeWithNext
	Right *TreeNodeWithNext
	Next  *TreeNodeWithNext
}

func connect(root *TreeNodeWithNext) *TreeNodeWithNext {
	rightEdge := make([]*TreeNodeWithNext, 12)

	var stitch func(depth int, node *TreeNodeWithNext)
	stitch = func(depth int, node *TreeNodeWithNext) {
		if node == nil {
			return
		}
		stitch(depth+1, node.Right)
		node.Next = rightEdge[depth]
		rightEdge[depth] = node
		stitch(depth+1, node.Left)
	}

	stitch(0, root)
	return root
}
