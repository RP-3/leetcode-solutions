// https://leetcode.com/problems/basic-calculator-ii/

package leetcode

/*
IDEA: Use a stack.
Given that we know the expression is valid parse adjacent numeric
characters into single numbers as we find them. We could do this
in multiple passes but this seems to be much faster.
*/
func calculate(s string) int {
	parsedNums := []int{}
	currNum, prevOp := 0, '+'
	s = s + "+" // add a tailing '+' so unary inputs aren't ignored

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' { // number
			currNum = currNum*10 + int(s[i]-'0')
		} else if s[i] == ' ' {
			continue
		} else {
			switch {
			case prevOp == '+':
				parsedNums = append(parsedNums, currNum)
			case prevOp == '-':
				parsedNums = append(parsedNums, -1*currNum)
			case prevOp == '*':
				left := parsedNums[len(parsedNums)-1]
				parsedNums[len(parsedNums)-1] = left * currNum
			case prevOp == '/':
				left := parsedNums[len(parsedNums)-1]
				parsedNums[len(parsedNums)-1] = left / currNum
			}
			prevOp = rune(s[i])
			currNum = 0
		}
	}

	result := 0
	for _, n := range parsedNums {
		result += n
	}
	return result
}
