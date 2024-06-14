package itinerary

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/route"
)

//go:generate mockgen -source=modifier.go -destination=mock/modifiermock.go
type Modifier interface {
	Info
	RemoveDroneStopFromFlight(index int, flight route.ISubRoute)
	RemoveMainStopFromRoute(index int)
	TryToInsertDroneDelivery(point gps.Point, calcCost func(Info) float64) bool
	TryToInsertIntoRoutes(point gps.Point, calcCost func(Info) float64) bool
}

type modifier struct {
	*info
}

func (m modifier) RemoveDroneStopFromFlight(index int, flight route.ISubRoute) {
	if flight.First() == flight.Last() {
		flight.StartingStop().RemoveStartingSubRoute(flight)
		flight.ReturningStop().RemoveReturningSubRoute(flight)
		m.removeFlightFromCompletedSubItineraryList(flight)
		return
	}

	flight.RemoveSubStop(index)
}

func (m modifier) RemoveMainStopFromRoute(index int) {
	m.route.RemoveMainStop(index)
}

func (m modifier) removeFlightFromCompletedSubItineraryList(flight route.ISubRoute) {
	newCompletedSubItineraryList := make([]SubItinerary, 0)
	for _, subItinerary := range m.completedSubItineraryList {
		if subItinerary.Flight != flight {
			newCompletedSubItineraryList = append(newCompletedSubItineraryList, subItinerary)
		}
	}
	m.completedSubItineraryList = newCompletedSubItineraryList
}

func (m modifier) TryToInsertDroneDelivery(point gps.Point, calcCost func(Info) float64) bool {
	lowestCost := -1.0
	lowestCostIndex := -1
	var lowestCostFlight route.ISubRoute

	for _, subItinerary := range m.completedSubItineraryList {
		if index, newCost, success := m.tryToInsertOnExistingFlight(subItinerary, point, calcCost); success {
			if lowestCostFlight == nil || newCost < lowestCost {
				lowestCost = newCost
				lowestCostIndex = index
				lowestCostFlight = subItinerary.Flight
			}
		}
	}

	if lowestCostFlight != nil {
		lowestCostFlight.InsertAt(lowestCostIndex, route.NewSubStop(point))
		return true
	}

	return false
}

func (m modifier) tryToInsertOnExistingFlight(subItinerary SubItinerary, point gps.Point, calcCost func(Info) float64) (int, float64, bool) {
	initialCost := calcCost(m)
	lowestCost := -1.0
	lowestCostIndex := -1

	for i := 0; i < subItinerary.Flight.Length(); i++ {
		flight := subItinerary.Flight
		flight.InsertAt(i, route.NewSubStop(point))
		if isValidFlight(subItinerary) {
			newCost := calcCost(m)
			if lowestCostIndex == -1 || newCost < lowestCost {
				lowestCost = newCost
				lowestCostIndex = i
			}
		}
		flight.RemoveSubStop(i)
	}

	if lowestCost == -1.0 {
		return 0, 0, false
	}

	addicionalCost := lowestCost - initialCost
	return lowestCostIndex, addicionalCost, true
}

func isValidFlight(subItinerary SubItinerary) bool {
	points := extractPointsFromFlight(subItinerary.Flight)
	pointsWithoutStartingPoint := points[1:]
	pointsWithoutStargingAndEndingPoint := points[1 : len(points)-1]
	subItinerary.Drone.Land(points[0])
	return subItinerary.Drone.CanReach(pointsWithoutStartingPoint...) &&
		subItinerary.Drone.Support(pointsWithoutStargingAndEndingPoint...)
}

func extractPointsFromFlight(flight route.ISubRoute) []gps.Point {
	points := []gps.Point{}
	points = append(points, flight.StartingStop().Point())
	iterator := flight.Iterator()
	points = append(points, iterator.Actual().Point())
	for iterator.HasNext() {
		points = append(points, iterator.Next().Point())
		iterator.GoToNext()
	}
	points = append(points, flight.ReturningStop().Point())
	return points
}

func (m modifier) TryToInsertIntoRoutes(point gps.Point, calcCost func(Info) float64) bool {
	lowestCost := -1.0
	lowestCostIndex := -1
	newCarStop := route.NewMainStop(point)

	for insertIndex := 1; insertIndex < m.route.Length()-1; insertIndex++ {
		m.route.InserAt(insertIndex, newCarStop)
		if isValidRoute(m) {
			newCost := calcCost(m)
			if lowestCostIndex == -1 || newCost < lowestCost {
				lowestCost = newCost
				lowestCostIndex = insertIndex
			}
		}
		m.route.RemoveMainStop(insertIndex)
	}

	if lowestCost == -1.0 {
		return false
	}

	m.route.InserAt(lowestCostIndex, newCarStop)
	return true
}

func isValidRoute(_ Itinerary) bool {
	return true
}
