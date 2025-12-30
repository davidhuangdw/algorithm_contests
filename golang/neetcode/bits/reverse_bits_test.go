package bits

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseBits(t *testing.T) {
	tests := []struct {
		name string
		n    uint32
		want uint32
	}{
		{
			name: "all zeros",
			n:    0b00000000000000000000000000000000,
			want: 0b00000000000000000000000000000000,
		},
		{
			name: "all ones",
			n:    0b11111111111111111111111111111111,
			want: 0b11111111111111111111111111111111,
		},
		{
			name: "single set bit at LSB",
			n:    0b00000000000000000000000000000001,
			want: 0b10000000000000000000000000000000,
		},
		{
			name: "single set bit at MSB",
			n:    0b10000000000000000000000000000000,
			want: 0b00000000000000000000000000000001,
		},
		{
			name: "alternating bits 1",
			n:    0b10101010101010101010101010101010,
			want: 0b01010101010101010101010101010101,
		},
		{
			name: "alternating bits 2",
			n:    0b01010101010101010101010101010101,
			want: 0b10101010101010101010101010101010,
		},
		{
			name: "leetcode example 1",
			n:    0b00000010100101000001111010011100,
			want: 0b00111001011110000010100101000000,
		},
		{
			name: "leetcode example 2",
			n:    0b11111111111111111111111111111101,
			want: 0b10111111111111111111111111111111,
		},
		{
			name: "pattern 1100",
			n:    0b11001100110011001100110011001100,
			want: 0b00110011001100110011001100110011,
		},
		{
			name: "pattern 0011",
			n:    0b00110011001100110011001100110011,
			want: 0b11001100110011001100110011001100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uint32(reverseBits(int(tt.n)))
			assert.Equal(t, tt.want, got, "reverseBits(%032b) = %032b, want %032b", tt.n, got, tt.want)
		})
	}
}
