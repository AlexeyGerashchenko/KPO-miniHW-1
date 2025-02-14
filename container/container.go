package container

import (
	"KPO-miniHW-1/services"
)

type Container struct {
	Zoo       *services.Zoo
	VetClinic services.VetClinic
}

func BuildContainer() *Container {
	vet := services.NewVetClinic()
	zoo := services.NewZoo(vet)
	return &Container{
		Zoo:       zoo,
		VetClinic: vet,
	}
}
