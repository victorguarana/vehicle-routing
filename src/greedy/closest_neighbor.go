package greedy

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/slc"
)

func ClosestNeighbor(itineraryList []itinerary.Itinerary, m gps.Map) {
	remaningClients := slc.Copy(m.Clients)
	for i := 0; len(remaningClients) > 0; i++ {
		itinerary := slc.CircularSelection(itineraryList, i)
		itineraryActualPosition := itinerary.ActualCarPoint()
		closestClient := gps.ClosestPoint(itineraryActualPosition, remaningClients)
		closestDepositFromClosestClient := gps.ClosestPoint(closestClient, m.Deposits)

		if itinerary.CarSupport(closestClient, closestDepositFromClosestClient) {
			itinerary.MoveCar(closestClient)
			remaningClients = slc.RemoveElement(remaningClients, closestClient)
			continue
		}

		closestDepositFromActualPosition := gps.ClosestPoint(itineraryActualPosition, m.Deposits)
		itinerary.MoveCar(closestDepositFromActualPosition)
	}

	finishItineraryOnClosestDeposits(itineraryList, m)
}
