package utils

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/structs"
)

func ChangeObjectName(animal *structs.Animal, name string) {
	animal.SetName(name)
}

func SpeakAnimal(randAnimal structs.Animal) {
	fmt.Printf("Animal says: %v\n", randAnimal.GetSound())
}
