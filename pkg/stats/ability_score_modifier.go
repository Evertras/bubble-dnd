package stats

import "fmt"

type AbilityModifier int

func (m AbilityModifier) String() string {
	if m >= 0 {
		return fmt.Sprintf("+%d", m)
	}

	return fmt.Sprintf("%d", m)
}
