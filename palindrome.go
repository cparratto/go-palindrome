package palindrome

import (
	"strings"

	stringutil "github.com/golang/example/stringutil"
)

//IsPalindrome returns true if string is
func IsPalindrome(s string) bool {
	var reversed string
	reversed = stringutil.Reverse(s)
	reversed = stripSpaces(reversed)
	s = stripSpaces(s)
	println("s = \t\t", s)
	println("reverse = \t", reversed)

	return s == reversed
}

func stripSpaces(s string) string {
	s = strings.ToLower(s)
	return strings.Replace(s, " ", "", -1)
}
