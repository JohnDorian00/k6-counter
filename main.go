package k6_counter

import (
	"github.com/JohnDorian00/k6-counter/counter"
	"go.k6.io/k6/js/modules"
)

type CounterModule struct{}

func (m *CounterModule) NewModuleInstance(vu modules.VU) interface{} {
	wrap := func(fn func() int64) func() int {
		callbackWrapper := vu.RegisterCallback()
		return func() int {
			var result int64
			callbackWrapper(func() error {
				result = fn()
				return nil
			})
			return int(result)
		}
	}

	return map[string]interface{}{
		"increment": wrap(counter.Increment),
		"decrement": wrap(counter.Decrement),
		"get":       wrap(counter.Get),
		"reset":     wrap(counter.Reset),
	}
}

func init() {
	modules.Register("k6/x/counter", &CounterModule{})
}
