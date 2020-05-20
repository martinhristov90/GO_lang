package word

// is s palindrome ?

func isPalindrome(s string) bool {
	for i := range s{
		if s[i] != s[len(s) - i - 1 ] {
			return false
		}
	}
	return true
}