package gender

import (
	"fmt"

	"github.com/jainSamkit/golang-design-patterns/enums"
)

//Male for male structure
type Male struct {
	FirstName string
	LastName  string
}

func (m *Male) GetName() string {
	name := fmt.Sprintf("%v %v", m.FirstName, m.LastName)
	return name
}

func (m *Male) GetGender() string {
	return enums.MALE
}
