package word

import "testing"

func TestPalindrom(t *testing.T){
	if !isPalindrome("detartrated") {
		t.Error(`isPalindrom("detartrated" = false`)
	}

	if !isPalindrome("kayak"){
		t.Error(`isPalindrom("kayak") = false`)
	}
}

func TestNonPalindrom(t *testing.T){
	if isPalindrome("palindrom"){
		t.Error(`isPalindrom("palindrom") = true`)
	}
}