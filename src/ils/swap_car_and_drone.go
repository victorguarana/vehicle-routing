package ils

import (
	"errors"

	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/measure"
)

func SwapCarAndDrone(modifier itinerary.Modifier, finder itinerary.Finder) error {
	worstDroneStopCost := finder.FindWorstDroneStop()
	swappableCarStopsOrdered := finder.FindWorstSwappableCarStopsOrdered()
	if len(swappableCarStopsOrdered) == 0 {
		return errors.New("No car stop was swappable")
	}

	modifier.RemoveDroneStopFromFlight(worstDroneStopCost.Index, worstDroneStopCost.Flight)
	if success := tryToInsertDroneDelivery(modifier, swappableCarStopsOrdered); !success {
		return errors.New("No car stop can be inserted into flight")
	}

	if err := modifier.InsertCarDelivery(worstDroneStopCost.Stop.Point(), measure.TotalDistance); err != nil {
		return errors.New("No drone stop can be inserted into route")
	}

	return nil
}

func tryToInsertDroneDelivery(modifier itinerary.Modifier, swappableCarStopsOrdered []itinerary.CarStopCost) bool {
	for _, carStopCost := range swappableCarStopsOrdered {
		if err := modifier.InsertDroneDelivery(carStopCost.Stop.Point(), measure.TotalDistance); err == nil {
			modifier.RemoveMainStopFromRoute(carStopCost.Index)
			return true
		}
	}

	return false
}
