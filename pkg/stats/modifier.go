package stats

import "fmt"

type Modifier int

func (m Modifier) String() string {
	if m >= 0 {
		return fmt.Sprintf("+%d", m)
	}

	return fmt.Sprintf("%d", m)
}
