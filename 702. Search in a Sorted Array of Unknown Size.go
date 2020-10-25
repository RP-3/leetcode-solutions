// https://leetcode.com/problems/search-in-a-sorted-array-of-unknown-size/

package leetcode

type ArrayReader struct{}

func (ar *ArrayReader) get(index int) int {}

func searchUnknownSize(reader ArrayReader, target int) int {
	l, r := 0, 2147483647
	for l <= r {
		mid := (l + r) / 2
		v := reader.get(mid)
		switch {
		case v == target:
			return mid
		case v < target:
			l = mid + 1
		case v > target:
			r = mid - 1
		}
	}

	return -1
}
