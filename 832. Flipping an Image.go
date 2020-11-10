// https://leetcode.com/problems/flipping-an-image/

package leetcode

func flipAndInvertImage(A [][]int) [][]int {
	result := make([][]int, len(A))

	for r, row := range A {
		result[r] = make([]int, len(row))
		rowLen := len(row)
		for c := rowLen - 1; c >= 0; c-- {
			offset := rowLen - c - 1
			result[r][offset] = (^row[c] & 1)
		}
	}

	return result
}
