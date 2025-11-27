package bootstrap

import (
	"github.com/alihdrii/go-hexagonal/pkg/di/container"
)

type Module interface {
	Register(mc container.MainContainerI)
}
