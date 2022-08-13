package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRollInExpectedRange(t *testing.T) {
	tests := []struct {
		name        string
		pool        Pool
		expectedMin int
		expectedMax int
	}{
		{
			name:        "1d4",
			pool:        PoolD4(1),
			expectedMin: 1,
			expectedMax: 4,
		},
		{
			name:        "5d6",
			pool:        PoolD6(5),
			expectedMin: 5,
			expectedMax: 30,
		},
		{
			name:        "3d8",
			pool:        PoolD8(3),
			expectedMin: 3,
			expectedMax: 24,
		},
		{
			name:        "2d10",
			pool:        PoolD10(2),
			expectedMin: 2,
			expectedMax: 20,
		},
		{
			name:        "4d12",
			pool:        PoolD12(4),
			expectedMin: 4,
			expectedMax: 48,
		},
		{
			name:        "2d20",
			pool:        PoolD20(2),
			expectedMin: 2,
			expectedMax: 40,
		},
	}

	const iterations = 10000

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i := 0; i < iterations; i++ {
				result := test.pool.Roll()

				assert.GreaterOrEqual(t, result.Total(), test.expectedMin)
				assert.LessOrEqual(t, result.Total(), test.expectedMax)

				assert.Len(t, result.RollResults(), test.pool.Count())

				if t.Failed() {
					t.FailNow()
				}
			}
		})
	}
}
