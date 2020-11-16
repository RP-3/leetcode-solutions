// https://leetcode.com/problems/longest-mountain-in-array/

package leetcode

func longestMountain(A []int) int {
	result, start := 0, len(A)
	asc := false

	for i := 1; i < len(A); i++ {
		switch {
		case A[i] > A[i-1] && asc: // was asc, still asc
			// do nothing
		case A[i] > A[i-1] && !asc: // was desc, now asc
			result = max(result, i-start)
			start = i - 1
			asc = true
		case A[i] < A[i-1] && asc: // was asc, now desc
			asc = false
			result = max(result, 1+i-start)
		case A[i] < A[i-1] && !asc: // was desc, still desc
			result = max(result, 1+i-start)
		default: // level ground. no mountain here
			start = len(A) // set start to OOB
			asc = false
		}
	}

	return result
}
