// https://leetcode.com/problems/complement-of-base-10-integer/

package leetcode

func bitwiseComplement(N int) int {

	if N == 0 {
		return 1 // because leetcode. I don't agree with this.
	}

	msbMask := N
	msbMask |= (msbMask >> 1)
	msbMask |= (msbMask >> 2)
	msbMask |= (msbMask >> 4)
	msbMask |= (msbMask >> 8)
	msbMask |= (msbMask >> 16)

	return (^N) & msbMask
}
