package bootstrap

import "github.com/alihdrii/go-hexagonal/pkg/di/container"

var modules []Module

func RegisterModule(m Module) {
	modules = append(modules, m)
}

func ApplyAll(container container.MainContainerI) {
	for _, m := range modules {
		m.Register(container)
	}
}
