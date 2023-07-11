package counter

import "sync/atomic"

// also can use sync.Mutex or sync.RWMutex
type сounter struct {
	value *atomic.Uint64
}

func New(init uint64) сounter {
	a := &atomic.Uint64{}
	a.Store(init)
	return сounter{value: a}
}

func (c *сounter) Inc() {
	c.Add(1)
}

func (c *сounter) Add(d uint64) {
	c.value.Add(d)
}

func (c *сounter) Value() uint64 {
	return c.value.Load()
}
