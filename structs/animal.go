package structs

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/interfaces"
)

type Animal struct {
	name   string
	weight int
	sound  string

	//strategy pattern.Also provide the flying type
	flyingType interfaces.Flys
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

func (animal *Animal) SetFlyingType(flyingtype interfaces.Flys) {
	animal.flyingType = flyingtype
}

func (animal *Animal) TrytoFly() string {
	return animal.flyingType.Fly()
}
