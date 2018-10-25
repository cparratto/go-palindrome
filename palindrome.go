package palindrome

import (
	"unicode"
)

// IsPalindrome returns true if string is
func IsPalindrome(s string) bool {
	var (
		runes  = []rune(s)
		front  = 0
		back   = len(runes) - 1
		middle = len(runes) / 2
	)

	for front < middle {
		if runes[front] == ' ' {
			front++
			continue
		}

		if runes[back] == ' ' {
			back--
			continue
		}

		if unicode.ToLower(runes[front]) != unicode.ToLower(runes[back]) {
			return false
		}

		front++
		back--
	}

	return true
}
