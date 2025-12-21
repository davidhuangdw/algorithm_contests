package arrays_hashing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolution_EncodeDecode(t *testing.T) {
	tests := []struct {
		strs []string
	}{
		{
			[]string{"hello", "world"},
		},
		{
			[]string{"test"},
		},
		{
			[]string{"a", "b", "c"},
		},
		{
			[]string{"", "", ""},
		},
		{
			[]string{"longer string with spaces", "short", ""},
		},
		{
			[]string{"special!@#$%^&*()", "characters"},
		},
		{
			[]string{"unicode: 你好世界", "🎉 emoji"},
		},
		{
			[]string{},
		},
	}

	solution := &Solution{}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("strs=%v", tt.strs), func(t *testing.T) {
			encoded := solution.Encode(tt.strs)
			decoded := solution.Decode(encoded)

			assert.Equal(t, decoded, tt.strs)
		})
	}
}
