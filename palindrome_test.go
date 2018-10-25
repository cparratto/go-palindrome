package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	palindromeTests := []struct {
		name      string
		candidate string
		valid     bool
	}{
		{candidate: "nurses run", valid: true},
		{candidate: "racecar", valid: true},
		{candidate: "notAplanidrome", valid: false},
		{candidate: "A dog a plan a canal pagoda", valid: true},
		{candidate: "A Toyotas a Toyota", valid: true},
		{candidate: "Ah, Satan sees Natasha", valid: true},

		

	}

	for _, pal := range palindromeTests {
		t.Run(pal.candidate, func(t *testing.T) {
			got := IsPalindrome(pal.candidate)
			want := pal.valid

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}
		})
	}

	got := IsPalindrome("nurses run")
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
