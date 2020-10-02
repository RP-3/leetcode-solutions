// https://leetcode.com/problems/combination-sum/

package leetcode

func combinationSum(candidates []int, target int) [][]int {
	result, wc := [][]int{}, []int{}

	var backtrack func(int, int)
	backtrack = func(i, remainder int) {
		switch {
		case i >= len(candidates) || remainder < 0:
			return
		case remainder == 0:
			dupe := make([]int, len(wc))
			copy(dupe, wc)
			result = append(result, dupe)
		default:
			wc = append(wc, candidates[i])
			backtrack(i, remainder-candidates[i])
			wc = wc[:len(wc)-1]
			backtrack(i+1, remainder)
		}
	}

	backtrack(0, target)
	return result
}
