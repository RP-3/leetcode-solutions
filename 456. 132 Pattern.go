// https://leetcode.com/problems/132-pattern/

package leetcode

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	// 1. find the min value (left inclusive) at every index
	mins := make([]int, len(nums))
	mins[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		mins[i] = min(mins[i-1], nums[i])
	}

	// 2. Maintain a set of elements > min at i and < current index
	rg := []int{nums[len(nums)-1]}

	// 3. Iterate from right to left, maintaining the invariant (2)
	for i := len(nums) - 2; i > 0; i-- {
		for len(rg) > 0 && rg[len(rg)-1] <= mins[i-1] {
			rg = rg[:len(rg)-1] // delete all items <= current min
		}
		if len(rg) == 0 || nums[i] < rg[len(rg)-1] {
			rg = append(rg, nums[i]) // add if less than current item
		}
		if len(rg) > 0 && nums[i] > rg[len(rg)-1] {
			return true // if at this point (2) > current item, return true!
		}
	}

	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
