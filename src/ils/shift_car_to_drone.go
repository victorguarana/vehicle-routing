package ils

import (
	"errors"
	"log"

	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/measure"
	"github.com/victorguarana/vehicle-routing/src/route"
	"github.com/victorguarana/vehicle-routing/src/slc"
)

type carStopCost struct {
	carStop route.IMainStop
	index   int
	cost    float64
}

func ShiftCarToDrone(modifier itinerary.Modifier) error {
	swappableCarStopsOrdered := findWorstSwappableCarStopsOrdered(modifier)
	if len(swappableCarStopsOrdered) == 0 {
		return errors.New("No car stop was swappable")
	}
	for _, carStopCost := range swappableCarStopsOrdered {
		if err := modifier.InsertDroneDelivery(carStopCost.carStop.Point(), measure.TotalDistance); err == nil {
			log.Printf("Car stop %v was shifted to drone by inserting into an existing flight", carStopCost.carStop)
			modifier.RemoveMainStopFromRoute(carStopCost.index)
			return nil
		}
	}

	return errors.New("No car stop can be inserted into flight")
}

// TODO: Try to move this function to modifier.go
func findWorstSwappableCarStopsOrdered(modifier itinerary.Modifier) []carStopCost {
	swappableCarStopsOrdered := []carStopCost{}
	iterator := modifier.RouteIterator()
	for iterator.HasNext() {
		iterator.GoToNext()
		actual := iterator.Actual()
		if !carStopIsSwappable(actual) {
			continue
		}
		carStopCost := carStopCost{
			carStop: actual,
			cost: gps.AdditionalDistancePassingThrough(
				iterator.Previous().Point(),
				actual.Point(),
				iterator.Next().Point(),
			),
			index: iterator.Index(),
		}
		swappableCarStopsOrdered = insertCarStopCostOrdered(swappableCarStopsOrdered, carStopCost)
	}
	return swappableCarStopsOrdered
}

func carStopIsSwappable(carStop route.IMainStop) bool {
	if carStop.IsWarehouse() {
		return false
	}
	if len(carStop.StartingSubRoutes()) > 0 {
		return false
	}
	if len(carStop.ReturningSubRoutes()) > 0 {
		return false
	}
	return true
}

func insertCarStopCostOrdered(carStopCosts []carStopCost, newCarStopCost carStopCost) []carStopCost {
	for i, csc := range carStopCosts {
		if newCarStopCost.cost > csc.cost {
			return slc.InsertAt(carStopCosts, newCarStopCost, i)
		}
	}
	return append(carStopCosts, newCarStopCost)
}
