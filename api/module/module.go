package module

import "go.uber.org/fx"

type Module interface {
	New() fx.Option
	Routes()
}

func All() []Module {
	return []Module{}
}
