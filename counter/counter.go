package counter

import (
	"sync/atomic"
)

// type CounterInterface interface {
// 	Increment(c Counter)
// 	Reset(c Counter)
// 	GetRequestCount(c Counter) *int64
// }

type Counter struct {
	count int64
}

func (c *Counter) Increment() {
	atomic.AddInt64(&c.count, 1)
}

func (c *Counter) Reset() {
	atomic.StoreInt64(&c.count, 0)
}

func (c *Counter) GetCurrentValue() *int64 {
	return &c.count
}

var RequestCounter = Counter{}
var ElapsedTimeInSecondsCounter = Counter{}
