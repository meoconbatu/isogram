package isogram

import (
	"strings"
	"unicode"
)

const testVersion = 1

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	for i, x := range word {
		if !unicode.IsLetter(x) {
			continue
		}
		for _, y := range word[i+1:] {
			if x == y {
				return false
			}
		}
	}
	return true
}
