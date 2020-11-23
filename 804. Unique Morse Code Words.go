// https://leetcode.com/problems/longest-substring-with-at-most-two-distinct-characters/

package leetcode

type void struct{}

var member void

func uniqueMorseRepresentations(words []string) int {
	// maps to the morse code for each character in the alphabet
	var code = [26]string{".-", "-...", "-.-.", "-..", ".",
		"..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
		"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-",
		".--", "-..-", "-.--", "--.."}

	dict := make(map[string]void)

	for _, word := range words {
		var codeWord strings.Builder
		for i := range word {
			index := int(word[i] - 97)
			codedChar := code[index]
			codeWord.WriteString(codedChar)
		}
		dict[codeWord.String()] = member
	}

	return len(dict)
}
