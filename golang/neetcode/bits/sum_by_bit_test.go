package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSum(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "both positive",
			a:    1,
			b:    2,
			want: 3,
		},
		{
			name: "both zero",
			a:    0,
			b:    0,
			want: 0,
		},
		{
			name: "one zero",
			a:    5,
			b:    0,
			want: 5,
		},
		{
			name: "negative and positive",
			a:    -1,
			b:    1,
			want: 0,
		},
		{
			name: "both negative",
			a:    -2,
			b:    -3,
			want: -5,
		},
		{
			name: "large numbers",
			a:    1000,
			b:    2000,
			want: 3000,
		},
		{
			name: "leetcode example 1",
			a:    1,
			b:    2,
			want: 3,
		},
		{
			name: "leetcode example 2",
			a:    2,
			b:    3,
			want: 5,
		},
		{
			name: "with carry",
			a:    3,
			b:    3,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSum(tt.a, tt.b)
			assert.Equal(t, tt.want, got, "getSum(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
		})
	}
}
