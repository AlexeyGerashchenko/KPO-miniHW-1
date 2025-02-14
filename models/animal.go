package models

var newAnimalID int = 1

type Animal struct {
	ID   int
	Name string
	Food int
}

func NewAnimal(name string, food int) Animal {
	newAnimal := Animal{
		ID:   newAnimalID,
		Name: name,
		Food: food,
	}
	newAnimalID++
	return newAnimal
}

func (animal Animal) GetNumber() int {
	return animal.ID
}

func (animal Animal) GetName() string {
	return animal.Name
}

func (animal Animal) GetFood() int {
	return animal.Food
}

type Herbo struct {
	Animal
	Kindness int
}

func NewHerbo(name string, food int, kindness int) Herbo {
	herbo := Herbo{
		Animal:   NewAnimal(name, food),
		Kindness: kindness,
	}
	return herbo
}

func (herbo Herbo) IsKind() bool {
	return herbo.Kindness > 5
}

type Predator struct {
	Animal
}

func NewPredator(name string, food int) Predator {
	predator := Predator{
		Animal: NewAnimal(name, food),
	}
	return predator
}

type Monkey struct {
	Animal
}

func NewMonkey(food int) Monkey {
	monkey := Monkey{
		Animal: NewAnimal("Monkey", food),
	}
	return monkey
}

type Rabbit struct {
	Herbo
}

func NewRabbit(food int, kindness int) Rabbit {
	rabbit := Rabbit{
		Herbo: NewHerbo("Rabbit", food, kindness),
	}
	return rabbit
}

type Tiger struct {
	Predator
}

func NewTiger(food int) Tiger {
	tiger := Tiger{
		Predator: NewPredator("Tiger", food),
	}
	return tiger
}

type Wolf struct {
	Predator
}

func NewWolf(food int) Wolf {
	wolf := Wolf{
		Predator: NewPredator("Wolf", food),
	}
	return wolf
}
