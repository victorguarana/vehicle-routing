package ils

import (
	"errors"
	"log"

	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/measure"
)

func SwapCarAndDrone(modifier itinerary.Modifier) error {
	worstDroneStop := findWorstDroneStop(modifier)
	swappableCarStopsOrdered := findWorstSwappableCarStopsOrdered(modifier)

	modifier.RemoveDroneStopFromFlight(worstDroneStop.index, worstDroneStop.flight)
	var err error
	for _, carStopCost := range swappableCarStopsOrdered {
		if err = modifier.InsertDroneDelivery(carStopCost.carStop.Point(), measure.TotalDistance); err == nil {
			log.Printf("Car stop %v was shifted to drone", carStopCost.carStop)
			modifier.RemoveMainStopFromRoute(carStopCost.index)
			break
		}
	}
	if err != nil {
		return errors.New("No car stop can be inserted into flight")
	}

	if err := modifier.InsertCarDelivery(worstDroneStop.droneStop.Point(), measure.TotalDistance); err != nil {
		return errors.New("No drone stop can be inserted into route")
	}

	return nil
}
