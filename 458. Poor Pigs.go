// https://leetcode.com/problems/poor-pigs/

package leetcode

import "math"

// log rule: log(b, x) = log(c, x) / log(c, b)
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	completeTrials := minutesToTest / minutesToDie
	b, ct := float64(buckets), float64(completeTrials)+1
	decimalResult := math.Log2(b) / math.Log2(ct)
	return int(math.Ceil(decimalResult))
}
