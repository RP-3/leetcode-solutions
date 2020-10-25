// https://leetcode.com/problems/bag-of-tokens/

package leetcode

import "sort"

func bagOfTokensScore(tokens []int, P int) int {
	sort.Ints(tokens)
	score, l, r := 0, 0, len(tokens)-1

	for l <= r {
		switch {
		// if we can get a point, do so
		case P >= tokens[l]:
			P -= tokens[l]
			score++
			l++
		// if losing a point allows us to get another point, do so
		case score >= 1 && P+tokens[r] >= tokens[l] && l < r:
			P += tokens[r]
			score--
			r--
		// there's nothing we can do. Stop.
		default:
			return score
		}
	}

	return score
}
