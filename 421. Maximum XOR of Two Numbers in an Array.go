// https://leetcode.com/problems/maximum-xor-of-two-numbers-in-an-array/

package leetcode

func findMaximumXOR(nums []int) int {
	result := 0
	root := trieNode{}

	for _, num := range nums {
		maxAgainstNum := root.maxAgainst(num)
		if maxAgainstNum > result {
			result = maxAgainstNum
		}
		root.insert(num)
	}

	return result
}

type trieNode struct {
	one  *trieNode
	zero *trieNode
}

func (t *trieNode) insert(val int) {
	mask := 1 << 31

	for mask > 0 { // assume 32-bit int is the max val
		bit := val & mask
		if bit > 0 {
			if t.one == nil {
				t.one = new(trieNode)
			}
			t = t.one
		} else {
			if t.zero == nil {
				t.zero = new(trieNode)
			}
			t = t.zero
		}
		mask = mask >> 1
	}
}

func (t *trieNode) maxAgainst(val int) int {
	mask := 1 << 31
	result := 0

	for mask > 0 && t != nil {
		bit := val & mask

		if bit > 0 { // bit is one
			if t.zero != nil { // use matching zero if possible
				result |= mask
				t = t.zero
			} else {
				t = t.one
			}
		} else { // bit is zero
			if t.one != nil { // use matching one if possible
				result |= mask
				t = t.one
			} else {
				t = t.zero
			}
		}

		mask = mask >> 1
	}

	return result
}
