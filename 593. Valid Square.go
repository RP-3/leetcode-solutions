// https://leetcode.com/problems/valid-square/

package leetcode

import "math"

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	distances, p := make(map[float64]int), [][]int{p1, p2, p3, p4}

	for i := 1; i < 4; i++ {
		for j := 0; j < i; j++ {
			if p[i][0] == p[j][0] && p[i][1] == p[j][1] {
				return false
			}
			distances[euclidian(p[i], p[j])]++
		}
	}

	return len(distances) == 2
}

func euclidian(p1, p2 []int) float64 {
	x1, y1 := float64(p1[0]), float64(p1[1])
	x2, y2 := float64(p2[0]), float64(p2[1])
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
