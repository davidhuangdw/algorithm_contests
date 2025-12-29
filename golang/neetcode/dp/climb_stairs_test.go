package dp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClimbStairs(t *testing.T) {
	t.Run("0 stairs", func(t *testing.T) {
		result := climbStairs(0)
		assert.Equal(t, 1, result) // 1 way: do nothing
	})

	t.Run("1 stair", func(t *testing.T) {
		result := climbStairs(1)
		assert.Equal(t, 1, result) // 1 way: 1 step
	})

	t.Run("2 stairs", func(t *testing.T) {
		result := climbStairs(2)
		assert.Equal(t, 2, result) // 2 ways: 1+1 or 2
	})

	t.Run("3 stairs", func(t *testing.T) {
		result := climbStairs(3)
		assert.Equal(t, 3, result) // 3 ways: 1+1+1, 1+2, 2+1
	})

	t.Run("4 stairs", func(t *testing.T) {
		result := climbStairs(4)
		assert.Equal(t, 5, result) // 5 ways: 1+1+1+1, 1+1+2, 1+2+1, 2+1+1, 2+2
	})

	t.Run("5 stairs", func(t *testing.T) {
		result := climbStairs(5)
		assert.Equal(t, 8, result) // Fibonacci sequence: 1,1,2,3,5,8
	})

	t.Run("10 stairs", func(t *testing.T) {
		result := climbStairs(10)
		assert.Equal(t, 89, result) // Fibonacci: 1,1,2,3,5,8,13,21,34,55,89
	})

	t.Run("large number", func(t *testing.T) {
		result := climbStairs(20)
		assert.Equal(t, 10946, result) // Verify algorithm handles larger inputs
	})
}
