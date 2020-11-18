// https://leetcode.com/problems/mirror-reflection/

package leetcode

func mirrorReflection(p int, q int) int {
	rightReflections, topReflections, y := 0, 0, q%p

	for y != 0 {
		y += q
		rightReflections++
		if y > p {
			topReflections++
		}
		y %= p
	}

	switch {
	case rightReflections%2 != 0:
		return 2
	case topReflections%2 == 0:
		return 1
	default:
		return 0
	}
}
