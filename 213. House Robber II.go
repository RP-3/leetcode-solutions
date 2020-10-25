// https://leetcode.com/problems/house-robber-ii/

package leetcode

func rob(nums []int) int {

	memo := make([]int, 2*len(nums))
	for i := range memo {
		memo[i] = -1
	}
	var m func(int, int) int
	m = func(i, canRobLast int) int {
		offset := len(nums) * canRobLast
		return offset + i
	}

	var r func(i, canRobLast int) int
	r = func(i, canRobLast int) int {
		memoIndex := m(i, canRobLast)

		switch {
		case i >= len(nums): // out of bounds. Can't rob anything
			return 0
		case i == len(nums)-1: // last house. Rob if we haven't robbed the first
			return canRobLast * nums[i]
		case memo[memoIndex] != -1: // Use the memo table if we've been here before
			return memo[memoIndex]
		case i == 0: // if we're starting off, decide whether or not to rob the first house
			memo[memoIndex] = max(
				nums[i]+r(i+2, 0), // rob this house
				0+r(i+1, 1),       // do not rob this house
			)
		default: // otherwise just decide whether or not to rob the current house
			memo[memoIndex] = max(
				nums[i]+r(i+2, canRobLast), // rob this house
				0+r(i+1, canRobLast),       // do not rob this house
			)
		}
		return memo[memoIndex]
	}

	return r(0, 0)
}
