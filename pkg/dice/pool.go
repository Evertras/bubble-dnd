package dice

import "fmt"

type Pool struct {
	kind  Kind
	count int
}

func (p *Pool) String() string {
	return fmt.Sprintf("%dd%d", p.count, p.kind)
}

func (p *Pool) Kind() Kind {
	return p.kind
}

func (p *Pool) Count() int {
	return p.count
}

func PoolD4(count int) Pool {
	return Pool{
		kind:  D4,
		count: count,
	}
}

func PoolD6(count int) Pool {
	return Pool{
		kind:  D6,
		count: count,
	}
}

func PoolD8(count int) Pool {
	return Pool{
		kind:  D8,
		count: count,
	}
}

func PoolD10(count int) Pool {
	return Pool{
		kind:  D10,
		count: count,
	}
}

func PoolD12(count int) Pool {
	return Pool{
		kind:  D12,
		count: count,
	}
}

func PoolD20(count int) Pool {
	return Pool{
		kind:  D20,
		count: count,
	}
}
