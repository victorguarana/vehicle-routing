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
	var success bool
	for _, carStopCost := range swappableCarStopsOrdered {
		if success := modifier.TryToInsertDroneDelivery(carStopCost.carStop.Point(), measure.TotalDistance); success {
			log.Printf("Car stop %v was shifted to drone", carStopCost.carStop)
			modifier.RemoveMainStopFromRoute(carStopCost.index)
			break
		}
	}
	if !success {
		return errors.New("No car stop can be inserted into flight")
	}

	if success := modifier.TryToInsertIntoRoutes(worstDroneStop.droneStop.Point(), measure.TotalDistance); !success {
		return errors.New("No drone stop can be inserted into route")
	}

	return nil
}
