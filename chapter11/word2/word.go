package word

import (
	"unicode"
)

func IsPalindrome(s string) bool {
	var letters []rune

	for _, r := range s { // range iterates over runes, not bytes
		if unicode.IsLetter(r) { //checks if the rune is an actual letter
			letters = append(letters, unicode.ToLower(r))
		}
	}

	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func IsPalindromeFast(s string) bool {
	var letters []rune

	for _, r := range s { // range iterates over runes, not bytes
		if unicode.IsLetter(r) { //checks if the rune is an actual letter
			letters = append(letters, unicode.ToLower(r))
		}
	}

	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func IsPalindromeUltraFast(s string) bool {
	
	letters := make([]rune, len(s))

	for _, r := range s { // range iterates over runes, not bytes
		if unicode.IsLetter(r) { //checks if the rune is an actual letter
			letters = append(letters, unicode.ToLower(r))
		}
	}

	n := len(letters) / 2
	for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}
