// https://leetcode.com/problems/merge-intervals/

package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Sort(byStartTime(intervals))
	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		prev, curr := result[len(result)-1], intervals[i]
		if prev[1] >= curr[0] {
			prev[1] = max(prev[1], curr[1])
		} else {
			result = append(result, curr)
		}
	}

	return result
}
