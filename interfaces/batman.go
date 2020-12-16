package interfaces

type Batman struct {
	name        string
	height      int
	weight      int
	responsible bool
}

func NewBatman() *Batman {
	return &Batman{
		name:        "Bale",
		height:      183,
		weight:      90,
		responsible: true,
	}
}

func (b Batman) GetWeight() int {
	return b.weight
}

func (b Batman) GetHeight() int {
	return b.height
}

func (b Batman) GetName() string {
	return b.name
}

func (b Batman) IsResponsible() bool {
	return b.responsible
}
