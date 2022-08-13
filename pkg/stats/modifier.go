package stats

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func modifierString[T constraints.Integer](val T) string {
	if val >= 0 {
		return fmt.Sprintf("+%d", val)
	}

	return fmt.Sprintf("%d", val)
}
