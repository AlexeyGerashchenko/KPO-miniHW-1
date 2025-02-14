package container

import (
	"KPO-miniHW-1/services"
)

type Container struct {
	Zoo       *services.Zoo
	VetClinic services.VetClinic
}
