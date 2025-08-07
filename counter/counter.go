package counter

import (
	"sync/atomic"
)

var GlobalCounter int64

func Increment() int64 {
	return atomic.AddInt64(&GlobalCounter, 1)
}

func Decrement() int64 {
	return atomic.AddInt64(&GlobalCounter, -1)
}

func Get() int64 {
	return atomic.LoadInt64(&GlobalCounter)
}

func Reset() int64 {
	atomic.StoreInt64(&GlobalCounter, 0)
	return 0
}
