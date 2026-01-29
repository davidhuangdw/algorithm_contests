package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModInv(t *testing.T) {
	tests := []struct {
		name string
		n    int
		p    int
		want []int
	}{
		{
			name: "n=1, p=5",
			n:    1,
			p:    5,
			want: []int{0, 1}, // inv[0] unused, inv[1] = 1
		},
		{
			name: "n=2, p=5",
			n:    2,
			p:    5,
			want: []int{0, 1, 3}, // 2*3=6 ≡ 1 mod 5
		},
		{
			name: "n=3, p=7",
			n:    3,
			p:    7,
			want: []int{0, 1, 4, 5}, // 2*4=8≡1, 3*5=15≡1 mod7
		},
		{
			name: "n=5, p=11",
			n:    5,
			p:    11,
			want: []int{0, 1, 6, 4, 3, 9}, // 2*6=12≡1, 3*4=12≡1, 4*3=12≡1, 5*9=45≡1 mod11
		},
		{
			name: "n=4, p=1000000007",
			n:    4,
			p:    1000000007,
			want: []int{0, 1, 500000004, 333333336, 250000002}, // Known values for large prime
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := modInv(tt.n, tt.p)
			assert.Equal(t, tt.want, got)

			// Verify that all inverses are correct: i * inv[i] ≡ 1 mod p
			for i := 1; i <= tt.n; i++ {
				if got[i] != tt.want[i] {
					continue // Skip if already failed equality check
				}
				assert.Equal(t, 1, (i*got[i])%tt.p, "inverse of %d modulo %d should satisfy (%d * %d) mod %d = 1", i, tt.p, i, got[i], tt.p)
			}
		})
	}
}

// TestModInvLarge tests with larger inputs to ensure efficiency
func TestModInvLarge(t *testing.T) {
	n := 1000
	p := 10007 // A prime number
	inv := modInv(n, p)

	assert.Len(t, inv, n+1)

	// Verify some random elements to ensure correctness
	for _, i := range []int{1, 10, 100, 500, 1000} {
		assert.Equal(t, 1, (i*inv[i])%p, "inverse verification failed for i=%d, p=%d", i, p)
	}
}
