// https://leetcode.com/problems/find-all-duplicates-in-an-array/

package leetcode

func findDuplicates(nums []int) []int {
	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n < 0 {
			n = -n
		}
		switch index := n - 1; {
		case nums[index] < 0:
			result = append(result, n)
		default:
			nums[index] = -nums[index]
		}
	}
	return result
}
