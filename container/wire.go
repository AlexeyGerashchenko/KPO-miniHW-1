//go:build wireinject
// +build wireinject

package container

import (
	"KPO-miniHW-1/services"

	"github.com/google/wire"
)

var ContainerSet = wire.NewSet(
	services.NewVetClinic,
	services.NewZoo,
	// Заполняем поля структуры Container по имени.
	wire.Struct(new(Container), "Zoo", "VetClinic"),
)

func InitializeContainer() *Container {
	wire.Build(ContainerSet)
	return &Container{}
}
