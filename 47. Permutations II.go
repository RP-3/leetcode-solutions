// https://leetcode.com/problems/permutations-ii/

package leetcode

import (
	"sort"
)

/*
IDEA: Enumerate all permutations in ascending order.
*/

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	// find the right-most occurence of nums in asc order
	var rightFoot = func() int {
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i] > nums[i-1] {
				return i - 1
			}
		}
		return -1
	}

	// reverse the elements in a slice between l, r (inclusive)
	var reverse = func(l, r int) {
		for l < r {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}

	// finds the next largest permutation of nums
	var permute = func() bool {
		foot := rightFoot()
		if foot == -1 {
			return false
		}

		// find the left-most element to the right of the foot
		// that is <= foot
		i := foot + 1
		for i < len(nums) && nums[i] > nums[foot] {
			i++
		}
		i--

		nums[foot], nums[i] = nums[i], nums[foot]
		reverse(foot+1, len(nums)-1)
		return true
	}

	result = append(result, clone(nums))
	for permute() {
		result = append(result, clone(nums))
	}
	return result
}

func clone(nums []int) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	return result
}
