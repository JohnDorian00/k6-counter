package k6_counter

import "go.k6.io/k6/js/modules"

const importPath = "k6/x/g-counter"

func init() {
	modules.Register(importPath, new(rootModule))
}
