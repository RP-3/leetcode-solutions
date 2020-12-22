// https://leetcode.com/problems/divide-chocolate/

package leetcode

func maximizeSweetness(s []int, K int) int {
	l, r := minIntSlice(s), sumIntSlice(s)
	result := l

	// private helper running in O(len(s))
	var allSlicesHaveAtLeast = func(mid int) bool {
		sum, cuts := 0, K
		for i, n := range s {
			sum += n
			if sum >= mid && i != len(s)-1 && cuts > 0 {
				cuts--
				sum = 0
			}
		}
		return cuts == 0 && sum >= mid
	}

	for l <= r {
		mid := (l + r) / 2
		if allSlicesHaveAtLeast(mid) {
			result = max(result, mid)
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return result
}
