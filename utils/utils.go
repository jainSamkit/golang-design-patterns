package utils

import (
	"github.com/jainSamkit/golang-design-patterns/structs"
)

func ChangeObjectName(animal *structs.Animal, name string) {
	animal.SetName(name)
}
