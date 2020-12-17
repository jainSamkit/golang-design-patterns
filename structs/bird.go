package structs

import (
	"github.com/jainSamkit/golang-design-patterns/interfaces"
)

type BirdAnimal struct {
	Animal
}

func NewBird() *BirdAnimal {

	bird := BirdAnimal{}
	bird.SetSound("Chirp")

	bird.SetFlyingType(interfaces.ItFlys{})

	return &bird
}
