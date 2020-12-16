//This file gives introduction to oops level concepts in golang.
//See the structs package that defines Animal ,CatAnimal and DogAnimal structs.
//DogAnimal and catAnimal are embedded with animal struct.
//There is no inheritance in go. Inheritance happens via struct embeddings.
//An embedded struct automatically contains all the methods of embeddings. For instance,dogAnimal struct has all the methods defined for animal struct.

package main

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/interfaces"

	"github.com/jainSamkit/golang-design-patterns/utils"

	"github.com/jainSamkit/golang-design-patterns/structs"
)

func main() {

	fido := structs.NewDog()
	fido.SetName("Fido")

	fmt.Printf("%v\n", fido.GetName())
	fido.DigHole()
	fido.SetWeight(-1)

	utils.ChangeObjectName(&fido.Animal, "Creole")
	fmt.Printf("%v\n", fido.GetName())

	//referring cat's animal embedding
	kitty := structs.NewCat().Animal
	//The above line is same as
	// Animal kitty = new Cat()

	doggy := structs.NewDog().Animal

	utils.SpeakAnimal(doggy)
	utils.SpeakAnimal(kitty)

	bale := interfaces.NewBatman()
	if utils.CheckISLiving(bale) {
		fmt.Printf("Its a living character\n")
	} else {
		fmt.Printf("Its a non living character\n")
	}
}
