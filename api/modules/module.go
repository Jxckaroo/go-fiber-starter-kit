package modules

import (
	"go.uber.org/fx"
)

type Module interface {
	New() fx.Option
	Routes()
}

func ToLoad() []Module {
	return []Module{}
}
