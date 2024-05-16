package greedy

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/itinerary"
)

func finishItineraryOnClosestDeposits(itineraryList []itinerary.Itinerary, m gps.Map) {
	for _, itinerary := range itineraryList {
		position := itinerary.ActualCarPoint()
		closestDeposit := gps.ClosestPoint(position, m.Deposits)
		itinerary.MoveCar(closestDeposit)
	}
}
