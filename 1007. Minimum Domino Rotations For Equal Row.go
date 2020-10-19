// https://leetcode.com/problems/minimum-domino-rotations-for-equal-row/

package leetcode

import "math"

func minDominoRotations(A []int, B []int) int {
	if len(A) == 0 {
		return 0
	}

	a1, a2 := countRotations(A[0], A, B)
	b1, b2 := countRotations(B[0], A, B)
	result := min2(a1, a2, b1, b2)
	if result == math.MaxInt64 {
		return -1
	}
	return result
}

func countRotations(val int, A []int, B []int) (int, int) {
	aCount, bCount := 0, 0
	for i := range A {
		if A[i] != val {
			if B[i] == val {
				aCount++
			} else {
				aCount = math.MaxInt64
				break
			}
		}
	}

	for i := range B {
		if B[i] != val {
			if A[i] == val {
				bCount++
			} else {
				bCount = math.MaxInt64
				break
			}
		}
	}

	return aCount, bCount
}

func min2(nums ...int) int {
	result := math.MaxInt64
	for _, n := range nums {
		if n < result {
			result = n
		}
	}
	return result
}
