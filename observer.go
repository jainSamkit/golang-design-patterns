//This module is going to talk about observer design pattern.
//This pattern is useful when a state change has to be relayed(observed in) to other different places.
//The publisher doesnt need to know anything about the observer state.

package main

import (
	"github.com/jainSamkit/golang-design-patterns/structs"
)

func main() {

	//Notice that none of the observer instance is called. Publisher doesnt know the nature of observers. It just knows that they
	// implement the interface Observer and hence are worthy of subscribers.

	//inits a publisher instance
	publisher := structs.StockPublisher{}

	//inits an observer instance
	structs.InitNewStockObserver(&publisher)

	//Modify publisher state.The observers consequently observe the states as well and prints out the prices.
	publisher.SetGooglPrice(24)
	publisher.SetAlphabetPrice(22)

	//Another instance of obserber is spawned.
	structs.InitNewStockObserver(&publisher)

	//Modify publisher state.changes observed by both the instances of observer. See the output screen.
	publisher.SetNetflixPrice(34)

}
