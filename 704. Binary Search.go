// https://leetcode.com/problems/binary-search/

package leetcode

func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (l + r) / 2
        switch {
        case nums[mid] < target:
            l = mid+1
        case nums[mid] > target:
            r = mid-1
        default:
            return mid
        }
    }
    return -1
}