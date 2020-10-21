// https://leetcode.com/problems/asteroid-collision

package leetcode

// Unnecessarily complicated divide and conquer approach.
// Could've just used a stack!
func asteroidCollision(asteroids []int) []int {
	length := len(asteroids)

	if length <= 1 { // if there's <=1 asteroid, no collisions possible
		return asteroids
	}

	// divide into two halves
	left := asteroidCollision(asteroids[0 : length/2])
	right := asteroidCollision(asteroids[length/2:])
	l, r := len(left)-1, 0 // points to the right edge of the left half, left end of the right half

	// INVARIANT: from this point onwards we can assume that the left and right halves
	// have no collisons within themselves. Merge the two halves together

	for l >= 0 && r < len(right) { // while there's something remaining in the left and right sides
		if left[l] > 0 && right[r] < 0 { // and there's a collision about to happen
			switch { // destroy one or both asteroids by incrementing the right pointer
			case left[l]+right[r] < 0:
				l--
			case left[l]+right[r] > 0:
				r++
			default:
				l--
				r++
			}
		} else {
			break
		}
	}
	// return a single collection with all collisions removed
	return append(left[:l+1], right[r:]...)
}
