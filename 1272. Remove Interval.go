// https://leetcode.com/problems/remove-interval/

package leetcode

func removeInterval(intervals [][]int, toBeRemoved []int) [][]int {
	result, nl, nr := make([][]int, 0), toBeRemoved[0], toBeRemoved[1]

	for _, interval := range intervals {
		l, r := interval[0], interval[1]
		switch {
		case l >= nl && r <= nr: // entirely contained
			break // ignore this interval
		case l < nl && r > nr: // entirely contains
			result = append(result, []int{l, nl}) // split into two
			result = append(result, []int{nr, r})
		case r >= nl && r <= nr: // straddles left end
			result = append(result, []int{l, min(r, nl)}) // truncate right end
		case l >= nl && l <= nr: // straddles right end
			result = append(result, []int{max(l, nr), r}) // truncate left end
		default:
			result = append(result, interval) // add the whole thing
		}
	}

	return result
}
