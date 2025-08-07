package k6_counter

import "go.k6.io/k6/js/modules"

type rootModule struct{}

func (*rootModule) NewModuleInstance(vu modules.VU) modules.Instance {
	return &module{vu}
}

type module struct {
	vu modules.VU
}

func (m *module) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]any{
			"increment": m.Increment,
			"decrement": m.Decrement,
			"get":       m.Get,
			"set":       m.Set,
			"add":       m.Add,
			"reset":     m.Reset,
		},
	}
}

var _ modules.Module = (*rootModule)(nil)
