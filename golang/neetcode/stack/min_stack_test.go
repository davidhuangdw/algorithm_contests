package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinStack(t *testing.T) {
	t.Run("ComplexSequence", func(t *testing.T) {
		//minStack := Constructor()
		minStack := ConstructorByDiff()

		// Complex sequence of operations
		operations := []struct {
			op          string
			value       int
			expectedTop int
			expectedMin int
		}{
			{"push", 10, 10, 10},
			{"push", 5, 5, 5},
			{"push", 15, 15, 5},
			{"push", 3, 3, 3},
			{"push", 6, 6, 3},
			{"pop", 0, 3, 3},  // pop 6
			{"pop", 0, 15, 5}, // pop 3
			{"push", 2, 2, 2},
			{"pop", 0, 15, 5}, // pop 2
			{"pop", 0, 5, 5},  // pop 15
			{"push", 1, 1, 1},
			{"pop", 0, 5, 5},   // pop 1
			{"pop", 0, 10, 10}, // pop 5
			{"pop", 0, 0, 0},   // pop 10
		}

		for _, op := range operations {
			switch op.op {
			case "push":
				minStack.Push(op.value)
			case "pop":
				minStack.Pop()
			}
			assert.Equal(t, op.expectedTop, minStack.Top(),
				"After %s %d, Top should be %d", op.op, op.value, op.expectedTop)
			assert.Equal(t, op.expectedMin, minStack.GetMin(),
				"After %s %d, GetMin should be %d", op.op, op.value, op.expectedMin)
		}
	})
}
