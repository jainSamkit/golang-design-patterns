package interfaces

type Living interface {
	GetName() string
	GetHeight() int
	GetWeight() int
	IsResponsible() bool
}
