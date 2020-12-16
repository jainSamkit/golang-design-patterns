package structs

import "fmt"

type DogAnimal struct {
	Animal
}

func NewDog() *DogAnimal {

	dog := DogAnimal{}
	dog.SetSound("Bark")

	return &dog
}

func (dog *DogAnimal) DigHole() {
	fmt.Println("Digging a hole!")
}
