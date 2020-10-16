// https://leetcode.com/problems/search-a-2d-matrix/

package leetcode

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows, cols := len(matrix), len(matrix[0])
	totalCells := rows * cols
	l, r := 0, totalCells-1

	for l <= r {
		mid := (l + r) / 2
		row, col := mid/cols, mid%cols
		val := matrix[row][col]

		switch {
		case val < target:
			l = mid + 1
		case val > target:
			r = mid - 1
		default:
			return true
		}
	}

	return false
}
