package factory

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/enums"
	"github.com/jainSamkit/golang-design-patterns/gender"
	"github.com/jainSamkit/golang-design-patterns/interfaces"
	"github.com/jainSamkit/golang-design-patterns/network"
)

func GetGenderObject(firstName, lastName string) (interfaces.Gender, error) {

	genderType, err := network.GetGender(firstName, lastName)
	if err != nil {
		fmt.Sprintln(err)
		return nil, err
	}
	if genderType == enums.MALE {
		maleObject := gender.Male{
			FirstName: firstName,
			LastName:  lastName,
		}
		gender := interfaces.Gender(&maleObject)
		return gender, nil
	} else {
		femaleObject := gender.Female{
			FirstName: firstName,
			LastName:  lastName,
		}
		gender := interfaces.Gender(&femaleObject)

		return gender, nil
	}
}
