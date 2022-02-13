package leetcode

import "strings"

func countVowels(word string) int64 {
	const VOWEL = "aeiou"
	var total int64
	n := len(word)
	for i, ch := range word {
		if strings.ContainsRune(VOWEL, ch) {
			total += int64(i+1) * int64(n-i)
		}
	}
	return total
}
