// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/

package leetcode

import (
	"strings"
)

func lengthOfLongestSubstringTwoDistinct(s string) int {
	result, l, r := 0, 0, 0
	counter := charCounter{}

	for r < len(s) {
		distinctCount := counter.add(s[r])
		for distinctCount > 2 {
			distinctCount = counter.remove(s[l])
			l++
		}

		result = max(result, 1+r-l)
		r++
	}

	return result
}

type charCounter struct {
	storage  [52]int
	distinct int
}

// add a char (case insensitive) to the counter
// and return the number of distinct chars
func (c *charCounter) add(char uint8) int {
	var index int
	if char >= 97 {
		index = int(char - 97)
	}else{
		index = int(char - 65) + 26
	}
	c.storage[index]++
	if c.storage[index] == 1 {
		c.distinct++
	}
	return c.distinct
}

// remove a char (case insensitive) from the counter
// and return the number of distinct chars
func (c *charCounter) remove(char uint8) int {
    var index int
	if char >= 97 {
		index = int(char - 97)
	}else{
		index = int(char - 65) + 26
	}
	c.storage[index]--
	if c.storage[index] == 0 {
		c.distinct--
	}
	return c.distinct
}
