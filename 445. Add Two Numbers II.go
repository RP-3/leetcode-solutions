// https://leetcode.com/problems/add-two-numbers-ii/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// convert the lists into stacks
	var stackify = func(node *ListNode) []int {
		result, p := []int{}, node
		for p != nil {
			result = append(result, p.Val)
			p = p.Next
		}
		return result
	}

	// use use nextVal as a handy wrapper to pop their values
	s1, s2, carry := stackify(l1), stackify(l2), 0
	var nextVal = func(isS1 bool) int {
		switch {
		case isS1 && len(s1) > 0:
			result := s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
			return result
		case !isS1 && len(s2) > 0:
			result := s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
			return result
		default:
			return 0
		}
	}

	var resultTail *ListNode
	for len(s1) != 0 || len(s2) != 0 {
		v3 := nextVal(true) + nextVal(false) + carry
		if v3 > 9 {
			carry = 1
			v3 = v3 - 10
		} else {
			carry = 0
		}
		newNode := &ListNode{Val: v3, Next: resultTail}
		resultTail = newNode
	}

	if carry == 1 {
		return &ListNode{Val: 1, Next: resultTail}
	}

	return resultTail
}
