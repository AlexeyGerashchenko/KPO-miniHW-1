package models

type IAlive interface {
	GetFood() int
}

type IInventory interface {
	GetNumber() int
	GetName() string
}

type IAnimal interface {
	IAlive
	IInventory
}
