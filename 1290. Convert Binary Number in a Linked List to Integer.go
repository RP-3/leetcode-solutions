// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	result := 0

	for head != nil {
		result <<= 1
		result |= head.Val
		head = head.Next
	}

	return result
}
