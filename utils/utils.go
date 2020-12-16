package utils

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/interfaces"

	"github.com/jainSamkit/golang-design-patterns/structs"
)

func ChangeObjectName(animal *structs.Animal, name string) {
	animal.SetName(name)
}

func SpeakAnimal(randAnimal structs.Animal) {
	fmt.Printf("Animal says: %v\n", randAnimal.GetSound())
}

func CheckISLiving(creature interfaces.Living) bool {
	if creature.GetHeight() > 0 && creature.GetWeight() > 0 && creature.GetName() != "" && creature.IsResponsible() {
		return true
	}
	return false
}
