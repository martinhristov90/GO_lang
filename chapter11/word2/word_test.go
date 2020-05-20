package word

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},  // non-palindrome
		{"desserts", false},    // semi-palindrome
		{"полиндром", false},   // non-polindrome in cyrillic
		{"Ядох и ходя.", true}, // polindrome in cyrillic
	}

	for _, test := range tests {
		if res := IsPalindrome(test.input); res != test.want {
			t.Errorf("tried IsPalindrome(%q) = %v", test.input, test.want)
		}
	}
}

func TestIsPalindromeFast(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false},  // non-palindrome
		{"desserts", false},    // semi-palindrome
		{"полиндром", false},   // non-polindrome in cyrillic
		{"Ядох и ходя.", true}, // polindrome in cyrillic
	}

	for _, test := range tests {
		if res := IsPalindromeFast(test.input); res != test.want {
			t.Errorf("tried IsPalindrome(%q) = %v", test.input, test.want)
		}
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("Ядох и ходя.")
	}
}

func BenchmarkIsPalindromeFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeFast("Ядох и ходя.")
	}
}

func BenchmarkIsPalindromeUltraFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindromeUltraFast("Ядох и ходя.")
	}
}
