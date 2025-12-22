package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		// Basic valid cases
		{"()", true},
		{"()[]{}", true},
		{"([])", true},
		{"{[]}", true},
		{"([{}])", true},
		{"((()))", true},
		{"[[[]]]", true},
		{"{{{{}}}}", true},
		{"([{}])[]", true},
		{"({[]}){}", true},

		// Basic invalid cases
		{"(", false},
		{")", false},
		{"()(", false},
		{"())", false},
		{"([)]", false},
		{"([)", false},
		{"({[}]", false},
		{"((())", false},
		{"[[[]]", false},
		{"{{{{}}}", false},

		// Edge cases
		{"", true},
		{"!@#", false},
		{"()a", false},
		{"a()", false},
		{"(a)", false},
		{"() ", false},
		{" () ", false},
		{"\n()\n", false},

		// Complex valid cases
		{"(([]){}){}", true},
		{"({})[({})]", true},
		{"[{()}]()", true},
		{"((((()))))", true},
		{"[[[[]]]]", true},
		{"{{{{{}}}}}", true},
		{"()()()()", true},
		{"[]{}()", true},
		{"([{}])([{}])", true},
		{"((()))[[[]]]{{{{}}}}", true},

		// Complex invalid cases
		{"(([]){}){", false},
		{"({})[({})", false},
		{"[{()}]()[", false},
		{"((((())))", false},
		{"[[[[]]]", false},
		{"{{{{{}}}}", false},
		{"()()()()(", false},
		{"[]{}(", false},
		{"([{}])([{}])", true},
		{"((()))[[[]]]{{{{}}}", false},

		// Mixed types
		{"({[]})", true},
		{"([{}])", true},
		{"{([])}", true},
		{"[({})]", true},
		{"({[}])", false},
		{"([{]}]", false},
		{"{(]}", false},
		{"[({]}", false},
		{"({[}]", false},
		{"([{})]", false},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			result := isValid(tt.s)
			assert.Equal(t, tt.want, result,
				"isValid(%q) = %v, want %v", tt.s, result, tt.want)
		})
	}
}
