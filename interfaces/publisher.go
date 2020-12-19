package interfaces

//Publisher interface defines all the methods for the puiblisher class. Has functionalities to register,unregister and notify observers.
type Publisher interface {
	Register(Observer)
	Unregister(Observer)
	NotifyObservers()
}
