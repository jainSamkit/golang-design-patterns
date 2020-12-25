package gender

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/enums"
)

//Female for female structure
type Female struct {
	FirstName string
	LastName  string
}

func (m *Female) GetName() string {
	name := fmt.Sprintf("%v %v", m.FirstName, m.LastName)
	return name
}

func (m *Female) GetGender() string {
	return enums.FEMALE
}
