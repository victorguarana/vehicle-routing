package ils

import (
	"errors"

	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/itinerary"
	"github.com/victorguarana/vehicle-routing/src/measure"
	"github.com/victorguarana/vehicle-routing/src/route"
)

type droneStopCost struct {
	index     int
	droneStop route.ISubStop
	flight    route.ISubRoute
	cost      float64
}

func ShiftDroneToCar(modifier itinerary.Modifier) error {
	worstDroneStop := findWorstDroneStop(modifier)
	err := modifier.InsertCarDelivery(worstDroneStop.droneStop.Point(), measure.TotalDistance)
	if err != nil {
		return errors.New("No drone stop can be inserted into route")
	}
	modifier.RemoveDroneStopFromFlight(worstDroneStop.index, worstDroneStop.flight)
	return nil
}

func findWorstDroneStop(modifier itinerary.Modifier) droneStopCost {
	worstDroneStop := droneStopCost{}
	subItineraryList := modifier.SubItineraryList()
	for _, subItinerary := range subItineraryList {
		worstDroneStopFromSubItinerary := findWorstSwappableDroneStopInFlight(subItinerary.Flight)
		if worstDroneStopFromSubItinerary.cost > worstDroneStop.cost {
			worstDroneStop = worstDroneStopFromSubItinerary
		}
	}

	return worstDroneStop
}

// TODO: Try to move this function to modifier.go
func findWorstSwappableDroneStopInFlight(flight route.ISubRoute) droneStopCost {
	iterator := flight.Iterator()
	iterator.GoToNext()
	// When the flight has only one stop
	if flight.First() == flight.Last() {
		return droneStopCost{
			index:     0,
			droneStop: flight.First(),
			flight:    flight,
			cost: gps.ManhattanDistanceBetweenPoints(
				flight.StartingStop().Point(),
				flight.First().Point(),
				flight.ReturningStop().Point(),
			),
		}
	}

	worstDroneStop := droneStopCost{}
	// When the flight has more than one stop
	previousPoint := flight.StartingStop().Point()
	for iterator.HasNext() {
		actualPoint := iterator.Actual().Point()
		nextPoint := iterator.Next().Point()
		cost := gps.ManhattanDistanceBetweenPoints(previousPoint, actualPoint, nextPoint)
		if cost > worstDroneStop.cost {
			worstDroneStop = droneStopCost{
				index:     iterator.Index(),
				droneStop: iterator.Actual(),
				flight:    flight,
				cost:      cost,
			}
		}
		previousPoint = actualPoint
		iterator.GoToNext()
	}

	actualPoint := iterator.Actual().Point()
	nextPoint := flight.ReturningStop().Point()
	cost := gps.ManhattanDistanceBetweenPoints(previousPoint, actualPoint, nextPoint)
	if cost > worstDroneStop.cost {
		worstDroneStop = droneStopCost{
			index:     iterator.Index(),
			droneStop: iterator.Actual(),
			flight:    flight,
			cost:      cost,
		}
	}

	return worstDroneStop
}
