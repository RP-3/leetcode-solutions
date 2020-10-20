// https://leetcode.com/problems/clone-graph/

package leetcode

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	clone := make(map[*Node]*Node)

	var dfs func(*Node) *Node
	dfs = func(n *Node) *Node {
		switch {
		case n == nil:
			return nil
		case clone[n] != nil:
			return clone[n]
		default:
			clone[n] = &Node{Val: n.Val}
			for _, neighbor := range n.Neighbors {
				clone[n].Neighbors = append(clone[n].Neighbors, dfs(neighbor))
			}
			return clone[n]
		}
	}

	return dfs(node)
}
