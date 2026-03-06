package main

import "testing"

func TestChineseRemainder(t *testing.T) {
	tests := []struct {
		name string
		ab   [][2]int
	}{
		{
			name: "simple case",
			ab:   [][2]int{{3, 2}, {5, 3}, {7, 2}},
		},
		{
			name: "two moduli",
			ab:   [][2]int{{5, 1}, {7, 3}},
		},
		//{
		//	name: "large numbers - overflow test",
		//	ab:   [][2]int{{99991, 99990}, {99989, 99988}, {99971, 99970}},
		//},
		{
			name: "single modulus",
			ab:   [][2]int{{13, 7}},
		},
		{
			name: "all zeros remainder",
			ab:   [][2]int{{3, 0}, {5, 0}, {7, 0}},
		},
		{
			name: "remainders equal to modulus - 1",
			ab:   [][2]int{{11, 10}, {13, 12}, {17, 16}},
		},
		{
			name: "coprime moduli",
			ab:   [][2]int{{97, 23}, {101, 45}, {103, 67}},
		},
		{
			name: "four moduli",
			ab:   [][2]int{{7, 3}, {11, 5}, {13, 10}, {17, 8}},
		},
		{
			name: "small moduli mixed remainders",
			ab:   [][2]int{{2, 1}, {3, 0}, {5, 4}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := chineseRemainder(tt.ab)

			// Verify result satisfies all congruences
			for _, p := range tt.ab {
				if result%p[0] != p[1] {
					t.Errorf("result=%d does not satisfy x ≡ %d (mod %d), got %d",
						result, p[1], p[0], result%p[0])
				}
			}
		})
	}
}
