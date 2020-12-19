package structs

import "github.com/jainSamkit/golang-design-patterns/interfaces"

//StockPublisher structure.Contains the list of observers at any time.
type StockPublisher struct {
	observers     []interfaces.Observer
	googlPrice    float64
	alphabetPrice float64
	netflixPrice  float64
}

//Register registers all the observers
func (pub *StockPublisher) Register(obs interfaces.Observer) {
	pub.observers = append(pub.observers, obs)
}

//UnRegister unregisters observer from the set of listeners
func (pub *StockPublisher) UnRegister(obs interfaces.Observer) {

	var todeleteIndex int
	for index, elem := range pub.observers {
		if elem == obs {
			todeleteIndex = index
			break
		}
	}

	pub.observers = append(pub.observers[:todeleteIndex], pub.observers[todeleteIndex+1:]...)
}

//SetGooglPrice sets google price and notify all the observers
func (pub *StockPublisher) SetGooglPrice(price float64) {
	pub.googlPrice = price

	for _, obs := range pub.observers {
		obs.Update(pub.netflixPrice, pub.googlPrice, pub.alphabetPrice)
	}
}

//SetNetflixPrice sets netflix price and notify all the observers
func (pub *StockPublisher) SetNetflixPrice(price float64) {
	pub.netflixPrice = price

	for _, obs := range pub.observers {
		obs.Update(pub.netflixPrice, pub.googlPrice, pub.alphabetPrice)
	}
}

//SetAlphabetPrice sets alphabet price and notify all the observers
func (pub *StockPublisher) SetAlphabetPrice(price float64) {
	pub.alphabetPrice = price

	for _, obs := range pub.observers {
		obs.Update(pub.netflixPrice, pub.googlPrice, pub.alphabetPrice)
	}
}
