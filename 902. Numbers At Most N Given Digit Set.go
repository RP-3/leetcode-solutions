// https://leetcode.com/problems/numbers-at-most-n-given-digit-set/

package leetcode

import (
	"math"
	"strconv"
)

/*
IDEA: build up the result from left to right, e.g., for inputs
- [1,3,5] n=4206
	- First count the 1-digit results
	- Then the 2-digit results... up to len(n)-1 = 3-digit results
	- You can do this easily because:
		- the number of n-digit results = #digits^n
		- any number with < len(n) digits must be < n
	- Now find the number of exactly len(n)-digit numbers that
	  can be constructred that are <= n. This is straight-up
	  backtrackng.
*/
func atMostNGivenDigitSet(digits []string, n int) int {
	// count the number of digits in n
	digitCount := 0
	for c := n; c > 0; c /= 10 {
		digitCount++
	}

	// convert the digits strings into more useful ints
	digitInts := make([]int, len(digits))
	for i, digit := range digits {
		parsedDigit, _ := strconv.ParseInt(digit, 10, 32)
		digitInts[i] = int(parsedDigit)
	}

	// count the number of distinct numbers we can make with
	// 1..[digitCount-1] digits.
	result := 0
	for i := 1; i < digitCount; i++ {
		result += pow(len(digits), i)
	}

	// use backtracking to find the count the number of distinct
	// numbers we can make with exactly digitCount digits.
	// This is feasible since we've written it cleverly to have
	// a branching factor of 1, and therefore a worst-case runtime
	// of O(digitCount * digits), which is at worst 9*9 = 81 (assuming
	// 32-bit ints)
	var count func(i int) int
	count = func(i int) int {
		if i == digitCount {
			return 1
		}

		// get the ith digit of n
		limitingDigit := (n / int(math.Pow10(digitCount-(i+1)))) % 10

		result := 0
		for _, d := range digitInts {
			switch {
			case d < limitingDigit:
				result += pow(len(digits), digitCount-(i+1))
			case d == limitingDigit:
				result += count(i + 1)
			default: // d > limitingDigit
				break // do nothing
			}
		}

		return result
	}

	return count(0) + result
}
