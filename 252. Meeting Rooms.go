// https://leetcode.com/problems/meeting-rooms/

package leetcode

import "sort"

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

type byStartTime [][]int

func (a byStartTime) Len() int           { return len(a) }
func (a byStartTime) Less(i, j int) bool { return a[i][0] < a[j][0] }
func (a byStartTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
