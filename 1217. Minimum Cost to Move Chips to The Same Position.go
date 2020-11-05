// https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/

package leetcode

func minCostToMoveChips(position []int) int {
	counts := make([]int, 2)
	for _, pos := range position {
		counts[pos%2]++
	}
	return min(counts[0], counts[1])
}
