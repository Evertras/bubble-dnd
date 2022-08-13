package stats

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModifierBase(t *testing.T) {
	tests := []struct {
		base             int
		expectedModifier int
	}{
		{
			base:             10,
			expectedModifier: 0,
		},
		{
			base:             11,
			expectedModifier: 0,
		},
		{
			base:             12,
			expectedModifier: 1,
		},
		{
			base:             16,
			expectedModifier: 3,
		},
		{
			base:             17,
			expectedModifier: 3,
		},
		{
			base:             21,
			expectedModifier: 5,
		},
		{
			base:             9,
			expectedModifier: -1,
		},
		{
			base:             8,
			expectedModifier: -1,
		},
		{
			base:             7,
			expectedModifier: -2,
		},
		{
			base:             6,
			expectedModifier: -2,
		},
		{
			base:             5,
			expectedModifier: -3,
		},
		{
			base:             3,
			expectedModifier: -4,
		},
		{
			base:             2,
			expectedModifier: -4,
		},
		{
			// Is there a minimum?
			base:             1,
			expectedModifier: -5,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%dIs%d", test.base, test.expectedModifier), func(t *testing.T) {
			s := New(Strength, test.base)

			assert.Equal(t, test.expectedModifier, s.Modifier())
		})
	}
}
