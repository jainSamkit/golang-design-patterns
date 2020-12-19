package structs

import (
	"fmt"
)

var observerID int

//StockObserver structure.
type StockObserver struct {
	id            int
	netflxPrice   float64
	googlPrice    float64
	alphabetPrice float64

	stockPublisher *StockPublisher
}

//InitNewStockObserver creates an stockobserver instance.It also subscribes to the publisher.
func InitNewStockObserver(publisher *StockPublisher) {

	observerID++
	observer := StockObserver{
		stockPublisher: publisher,
		id:             observerID,
	}

	//registers with the publisher
	publisher.Register(&observer)

}

//Update updates and print the updates prices. The update function is a callback to the state change iof publisher.
func (obs *StockObserver) Update(netflixPrice float64, googlPrice float64, alphabetPrice float64) {
	obs.netflxPrice = netflixPrice
	obs.googlPrice = googlPrice
	obs.alphabetPrice = alphabetPrice

	obs.printPrices()
}

//PrintPrices prints all the prices
func (obs *StockObserver) printPrices() {
	fmt.Printf("ID: %v GooglePrice: %v NetflixPrice: %v AlphabetPrice: %v\n", obs.id, obs.googlPrice, obs.netflxPrice, obs.alphabetPrice)
}
