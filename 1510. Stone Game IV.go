// https://leetcode.com/problems/stone-game-iv/

package leetcode

import "math"

func winnerSquareGame(n int) bool {
	has, cached := newBitset(n+1), newBitset(n+1)

	var currentPlayerWins func(int) bool
	currentPlayerWins = func(n int) bool {
		switch {
		case n == 0:
			return false
		case has.Get(n):
			return cached.Get(n)
		default:
			has.Set(n)
			root := int(math.Sqrt(float64(n)))
			for i := 1; i <= root; i++ {
				if !currentPlayerWins(n - i*i) {
					cached.Set(n)
					return true
				}
			}
			return false
		}
	}

	return currentPlayerWins(n)
}

// use a nifty bitset for speed
const wordSize = 64

type bitset struct {
	data []uint64
}

func newBitset(size int) *bitset {
	return &bitset{data: make([]uint64, (size/wordSize)+1)}
}

func (b *bitset) Get(x int) bool {
	return b.data[x/wordSize]&(1<<(x%wordSize)) != 0
}

func (b *bitset) Set(x int) {
	b.data[x/wordSize] |= (1 << (x % wordSize))
}
