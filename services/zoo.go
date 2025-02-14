package services

import (
	"KPO-miniHW-1/models"
	"fmt"
)

type Zoo struct {
	animals   []models.IAnimal
	things    []models.IInventory
	vetClinic VetClinic
}

func NewZoo(vetClinic VetClinic) *Zoo {
	return &Zoo{
		animals:   make([]models.IAnimal, 0),
		things:    make([]models.IInventory, 0),
		vetClinic: vetClinic,
	}
}

func (z *Zoo) AddAnimal(a models.IAnimal) {
	fmt.Println("Проверка здоровья животного...")
	healthy := z.vetClinic.CheckHealth()
	if healthy {
		z.animals = append(z.animals, a)
		fmt.Println("Животное принято в зоопарк.")
	} else {
		fmt.Println("Животное не принято в зоопарк из-за плохого здоровья.")
	}
}

func (z *Zoo) GetAnimals() []models.IAnimal {
	return z.animals
}

func (z *Zoo) GetTotalFoodConsumption() int {
	total := 0
	for _, a := range z.animals {
		total += a.GetFood()
	}
	return total
}

func (z *Zoo) GetContactZooEligibleAnimals() []models.IAnimal {
	var eligible []models.IAnimal
	for _, a := range z.animals {
		if ce, ok := a.(interface{ IsContactZooEligible() bool }); ok {
			if ce.IsContactZooEligible() {
				eligible = append(eligible, a)
			}
		}
	}
	return eligible
}

func (z *Zoo) PrintInventory() {
	fmt.Println("Инвентарный учёт животных:")
	for _, a := range z.animals {
		fmt.Printf("Наименование: %s, Инв. номер: %d\n", a.GetName(), a.GetNumber())
	}
	fmt.Println("Инвентарный учёт вещей:")
	for _, t := range z.things {
		fmt.Printf("Наименование: %s, Инв. номер: %d\n", t.GetName(), t.GetNumber())
	}
}

func (z *Zoo) AddThing(t models.IInventory) {
	z.things = append(z.things, t)
}
