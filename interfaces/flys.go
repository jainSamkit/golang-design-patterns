package interfaces

//The code below introduces an interface and two strategies on how it is implemented.

type Flys interface {
	Fly() string
}

type ItFlys struct {
}

func (itFlys ItFlys) Fly() string {

	return "Flying high"

}

type CantFly struct {
}

func (cantFly CantFly) Fly() string {

	return "Can't fly."
}
