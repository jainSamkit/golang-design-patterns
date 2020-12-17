//This file introduces strategy pattern.
//Strategy pattern is used when all of the subclasses show similar kinds of behaviour.(Fly or Cant fly).
//It is also used to dynamically change the behaviour of subclasses at run time.

//Interfaces package defines a flys(top behaviour/concept) interface. Two structs or two kinds of behaviours
//of flying object is created that implement fly interface.

//dog and Bird each set their behaviour.

//Refer to the code for more details.

package main

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/interfaces"

	"github.com/jainSamkit/golang-design-patterns/structs"
)

func main() {

	sparky := structs.NewDog().Animal
	tweety := structs.NewBird().Animal

	fmt.Printf("Dog: %v\n", sparky.TrytoFly())
	fmt.Printf("Bird: %v\n", tweety.TrytoFly())

	//dynamically changing flying character of dog
	sparky.SetFlyingType(interfaces.ItFlys{})
	fmt.Printf("Dog: %v\n", sparky.TrytoFly())
}
