package ils

import (
	"errors"

	"github.com/victorguarana/vehicle-routing/internal/itinerary"
	"github.com/victorguarana/vehicle-routing/internal/measure"
)

func ShiftDroneToCar(modifier itinerary.Modifier, finder itinerary.Finder) error {
	worstDroneStopCost := finder.FindWorstDroneStop()
	err := modifier.InsertCarDelivery(worstDroneStopCost.Stop.Point(), measure.TotalDistance)
	if err != nil {
		return errors.New("No drone stop can be inserted into route")
	}
	modifier.RemoveDroneStopFromFlight(worstDroneStopCost.Index, worstDroneStopCost.Flight)
	return nil
}
