// https://leetcode.com/problems/subarray-product-less-than-k/

package leetcode

func numSubarrayProductLessThanK(nums []int, k int) int {
	result, l, prod := 0, 0, 1
	for r := 0; r < len(nums); r++ {
		prod *= nums[r]
		for l < r && prod >= k {
			prod /= nums[l]
			l++
		}
		if prod < k {
			result += (r - l + 1)
		}
	}
	return result
}
