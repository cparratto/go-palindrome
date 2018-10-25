package palindrome

import (
	"strings"
)

//IsPalindrome returns true if string is
func IsPalindrome(s string) bool {
	r := []rune(strings.ToLower(strings.Replace(s, " ", "", -1)))

	for i := 0; i < len(r)/2; i++ {
		if !matchesAt(i, r) {
			return false
		}
	}

	return true
}

func matchesAt(pos int, pal []rune) bool {
	e := len(pal) - 1 - pos
	return pal[pos] == pal[e]
}
