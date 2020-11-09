// https://leetcode.com/problems/two-sum-less-than-k/

package leetcode

import "sort"

func twoSumLessThanK(A []int, K int) int {
	sort.Ints(A)
	if len(A) < 2 || A[0]+A[1] > K {
		return -1
	}

	l, r, result := 0, len(A)-1, A[0]+A[1]

	for l < r {
		sum := A[l] + A[r]
		switch {
		case sum >= K:
			r--
		case sum == K-1:
			return sum
		default: // sum < K
			result = max(result, sum)
			l++
		}
	}

	return result
}
