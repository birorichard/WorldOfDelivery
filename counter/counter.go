package counter

import "sync/atomic"

var requestCount int64

func Increment() {
	atomic.AddInt64(&requestCount, 1)
}

func Reset() {
	atomic.StoreInt64(&requestCount, 0)
}

func GetRequestCount() int64 {
	return requestCount
}
