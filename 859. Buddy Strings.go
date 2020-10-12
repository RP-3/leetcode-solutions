// https://leetcode.com/problems/buddy-strings/

package leetcode

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) || len(A) < 2 {
		return false
	}

	d := [][]byte{}
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			d = append(d, []byte{A[i], B[i]})
		}
	}

	if len(d) > 2 || len(d) == 1 {
		return false
	}

	if len(d) == 2 {
		return d[0][0] == d[1][1] && d[0][1] == d[1][0]
	}

	counts := make([]int, 26)
	for i := 0; i < len(A); i++ {
		c := int(rune(A[i])) - 97
		counts[c]++
		if counts[c] > 1 {
			return true
		}
	}

	return false
}
