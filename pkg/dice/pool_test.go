package dice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoolStrings(t *testing.T) {
	tests := []struct {
		name        string
		pool        Pool
		expectedMin int
		expectedMax int
	}{
		{
			name: "1d4",
			pool: PoolD4(1),
		},
		{
			name: "5d4",
			pool: PoolD4(5),
		},
		{
			name: "3d6",
			pool: PoolD6(3),
		},
		{
			name: "2d8",
			pool: PoolD8(2),
		},
		{
			name: "8d10",
			pool: PoolD10(8),
		},
		{
			name: "9d12",
			pool: PoolD12(9),
		},
		{
			name: "10d20",
			pool: PoolD20(10),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			str := test.pool.String()

			assert.Equal(t, test.name, str)
		})
	}
}
