// https://leetcode.com/problems/meeting-rooms/

package leetcode

import (
	"sort"
)

func canAttendMeetings(intervals [][]int) bool {
	sort.Sort(byStartTime(intervals))
	for i := 1; i < len(intervals); i++ {
		curr, prev := intervals[i], intervals[i-1]
		if prev[1] > curr[0] {
			return false
		}
	}
	return true
}
