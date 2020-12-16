package structs

import (
	"fmt"
)

type Animal struct {
	name   string
	weight int
	sound  string
}

func (animal *Animal) SetName(name string) {
	animal.name = name
}

func (animal *Animal) GetName() string {
	return animal.name
}

func (animal *Animal) SetSound(sound string) {
	animal.sound = sound
}

func (animal *Animal) GetSound() string {
	return animal.sound
}
func (animal *Animal) GetWeight() int {
	return animal.weight
}

func (animal *Animal) SetWeight(weight int) {
	if weight > 0 {
		animal.weight = weight
	} else {
		fmt.Printf("Weight should be greater than zero\n")
	}
}

func (animal *Animal) ChangeObjectName(dog *DogAnimal) {
	dog.SetName("Fido1")
}
