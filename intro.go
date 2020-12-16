//This file gives introduction to oops level concepts in golang

package main

import (
	"fmt"

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
}
