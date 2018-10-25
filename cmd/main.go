package main

import (
	"os"

	"github.com/mspeer383/go-palindrome"
)

func main() {
	
	candidate := os.Args[1]
	myp := palindrome.IsPalindrome(candidate)

	if myp {
		println("Found Palindrome ",candidate)
	} else {
		println("Did NOT find Palidrome", candidate)
	}
}
