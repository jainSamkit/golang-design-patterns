package structs

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/interfaces"
)

type DogAnimal struct {
	Animal
}

func NewDog() *DogAnimal {

	dog := DogAnimal{}
	dog.SetSound("Bark")

	dog.SetFlyingType(interfaces.CantFly{})

	return &dog
}

func (dog *DogAnimal) DigHole() {
	fmt.Println("Digging a hole!")
}
