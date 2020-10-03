// https://leetcode.com/problems/k-diff-pairs-in-an-array/

package leetcode

import "sort"

func findPairs(nums []int, k int) int {
	sort.Ints(nums) // sort ascending
	result, l, r := 0, 0, 1

	for r < len(nums) {
		diff := nums[r] - nums[l]

		switch {
		case l == r:
			r++ // this is not a valid pair
		case diff < k:
			r++ // make the diff larger
		case diff > k:
			l++ // make the diff smaller
		default: // diff == k
			result++ // record this result
			r++      // shift r until we guarantee no duplicates
			for r < len(nums) && nums[r] == nums[r-1] {
				r++
			}
		}
	}

	return result
}
