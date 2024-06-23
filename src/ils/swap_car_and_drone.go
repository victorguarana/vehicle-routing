package ils

import (
	"errors"
	"log"

	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/measure"
)

func SwapCarAndDrone(modifier itinerary.Modifier, finder itinerary.Finder) error {
	worstDroneStopCost := finder.FindWorstDroneStop()
	swappableCarStopsOrdered := finder.FindWorstSwappableCarStopsOrdered()

	modifier.RemoveDroneStopFromFlight(worstDroneStopCost.Index, worstDroneStopCost.Flight)
	var err error
	for _, carStopCost := range swappableCarStopsOrdered {
		if err = modifier.InsertDroneDelivery(carStopCost.Stop.Point(), measure.TotalDistance); err == nil {
			log.Printf("Car stop %v was shifted to drone", carStopCost.Stop)
			modifier.RemoveMainStopFromRoute(carStopCost.Index)
			break
		}
	}
	if err != nil {
		return errors.New("No car stop can be inserted into flight")
	}

	if err := modifier.InsertCarDelivery(worstDroneStopCost.Stop.Point(), measure.TotalDistance); err != nil {
		return errors.New("No drone stop can be inserted into route")
	}

	return nil
}
