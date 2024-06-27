package ils

import (
	"errors"
	"log"

	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/measure"
)

func ShiftCarToDrone(modifier itinerary.Modifier, finder itinerary.Finder) error {
	swappableCarStopsOrdered := finder.FindWorstSwappableCarStopsOrdered()
	if len(swappableCarStopsOrdered) == 0 {
		return errors.New("No car stop was swappable")
	}
	for _, carStopCost := range swappableCarStopsOrdered {
		if err := modifier.InsertDroneDelivery(carStopCost.Stop.Point(), measure.TotalDistance); err == nil {
			log.Printf("Car stop %v was shifted to drone by inserting into an existing flight", carStopCost.Stop)
			modifier.RemoveMainStopFromRoute(carStopCost.Index)
			return nil
		}
	}

	return errors.New("No car stop can be inserted into flight")
}
