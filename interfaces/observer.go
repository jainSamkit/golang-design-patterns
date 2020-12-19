package interfaces

//Observer interface has only one method update that the publisher will call to notify it. Notice that the update function is capitalised.
type Observer interface {
	Update(float64, float64, float64)
}
