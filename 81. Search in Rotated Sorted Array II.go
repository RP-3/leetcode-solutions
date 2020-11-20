// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/

package leetcode

func search(nums []int, target int) bool {

	var inner func(l, r int) bool
	inner = func(l, r int) bool {
		for l <= r {
			mid := (l + r) / 2

			switch {
			case nums[mid] == target:
				return true
			case nums[mid] == nums[r]: // pivot could be left or right. search both left and right
				return inner(l, mid-1) || inner(mid+1, r)
			case nums[mid] > nums[r]: // pivot between mid and r
				if inner(l, mid-1) { // normal binsearch on left
					return true
				}
				l = mid + 1 // continue searching for pivot on right
			default: // nums[mid] < nums[r]; pivot between l and mid
				if inner(mid+1, r) { // normal binsearch on right
					return true
				}
				r = mid - 1 // continue searching for pivot on left
			}
		}

		return false
	}

	return inner(0, len(nums)-1)
}
