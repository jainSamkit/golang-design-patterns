//This module demontrates factory pattern in golang.package patterns

//Factory pattern is useful when we need dynamic class instances(i.e creating classes at runtime).
//The above classes in most cases have a common parent class.
//Encapsulation of class building

//The program asks user to enter a his/her name and returns class based on his gender.Male for Man for males and female for females

package main

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/factory"
)

func main() {
	var firstName string
	var lastName string
	fmt.Println("Enter first name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter second name: ")
	fmt.Scanln(&lastName)

	genderObject, _ := factory.GetGenderObject(firstName, lastName)
	if genderObject != nil {
		fmt.Println("Name: ", genderObject.GetName())
		fmt.Println("Gender: ", genderObject.GetGender())
	}
}
