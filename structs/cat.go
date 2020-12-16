package structs

type CatAnimal struct {
	Animal
}

func NewCat() *CatAnimal {

	cat := CatAnimal{}
	cat.SetSound("Meow")

	return &cat
}
