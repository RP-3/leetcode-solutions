// https://leetcode.com/problems/repeated-dna-sequences/

package leetcode

import (
	"fmt"
	"strings"
)

const tides = "ACGT"

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	result, seen, hash := []string{}, make(map[int]int), 0
	for i := 0; i < 10; i++ {
		hash = hashChar(hash, s[i])
	}
	seen[hash]++

	for i := 10; i < len(s); i++ {
		hash = hashChar(hash, s[i])
		seen[hash]++
	}

	for hash, count := range seen {
		if count > 1 {
			result = append(result, unhash(hash))
		}
	}

	return result
}

func unhash(hash int) string {
	var b strings.Builder
	b.Grow(10)

	for i := 0; i < 10; i++ {
		mask := 3 << (18 - 2*i)
		charCode := (mask & hash) >> (18 - 2*i)
		fmt.Fprintf(&b, string(tides[charCode]))
	}

	return b.String()
}

func hashChar(hash int, char byte) int {
	hash <<= 2
	hash += strings.Index(tides, string(char))
	return hash & 0xFFFFF // 20 bits
}
