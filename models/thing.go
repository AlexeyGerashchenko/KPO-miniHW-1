package models

var nextThingID int = 1

type Thing struct {
	ID   int
	Name string
}

func NewThing(name string) Thing {
	t := Thing{
		ID:   nextThingID,
		Name: name,
	}
	nextThingID++
	return t
}

func (t Thing) GetNumber() int {
	return t.ID
}

func (t Thing) GetName() string {
	return t.Name
}

type Table struct {
	Thing
}

func NewTable() Table {
	return Table{
		Thing: NewThing("Table"),
	}
}

type Computer struct {
	Thing
}

func NewComputer() Computer {
	return Computer{
		Thing: NewThing("Computer"),
	}
}
