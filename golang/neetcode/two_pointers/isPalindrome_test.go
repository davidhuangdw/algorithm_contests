package two_pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
		{" ", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"0P", false},
		{"racecar", true},
		{"Able was I ere I saw Elba", true},
		{"Madam, I'm Adam", true},
		{"Was it a car or a cat I saw?", true},
		{"No 'x' in Nixon", true},
		{"hello", false},
		{"12321", true},
		{"123a321", true},
		{"A1b2c3c2b1a", true},
		{"A1b2c3d2b1a", false},
		{"!!!", true},
		{"a!!!a", true},
		{"a!!!b", false},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.input), 
				"isPalindrome(%q) should be %v", tt.input, tt.want)
		})
	}
}

func TestIsPalindrome_EdgeCases(t *testing.T) {
	assert.True(t, isPalindrome(""), "empty string should be palindrome")
	assert.True(t, isPalindrome("a"), "single character should be palindrome")
	assert.True(t, isPalindrome("A"), "single uppercase character should be palindrome")
	assert.True(t, isPalindrome("1"), "single digit should be palindrome")
	assert.True(t, isPalindrome(".,!@#$"), "only non-alphanumeric should be palindrome")
	assert.False(t, isPalindrome("ab"), "two different characters should not be palindrome")
	assert.True(t, isPalindrome("aa"), "two same characters should be palindrome")
}