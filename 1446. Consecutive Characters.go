// https://leetcode.com/problems/consecutive-characters/

package leetcode

func maxPower(s string) int {
	result := 1
	prevChar := s[0]
	currentPower := 1
	for i := 1; i < len(s); i++ {
		if s[i] == prevChar {
			currentPower++
			if currentPower > result {
				result = currentPower
			}
		} else {
			prevChar = s[i]
			currentPower = 1
		}
	}
	return result
}
