package itinerary

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/routes"
	"github.com/victorguarana/go-vehicle-route/src/slc"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

var flightFactory = routes.NewSubRoute

// TODO: Hide its implementation from other packages tests
// Avoid 'var mockedDrone1 = itinerary.DroneNumber(1)'
type DroneNumber int

type Itinerary interface {
	ActualCarPoint() gps.Point
	CarSupport(nextPoints ...gps.Point) bool
	DroneCanReach(droneNumber DroneNumber, nextPoints ...gps.Point) bool
	DroneNumbers() []DroneNumber
	DroneIsFlying(droneNumber DroneNumber) bool
	DroneSupport(droneNumber DroneNumber, nextPoints ...gps.Point) bool
	LandAllDrones(landingStop routes.IMainStop)
	LandDrone(droneNumber DroneNumber, destination routes.IMainStop)
	MoveCar(destination gps.Point)
	MoveDrone(droneNumber DroneNumber, destination gps.Point)
	RemoveMainStopFromRoute(index int)
	RouteIterator() slc.Iterator[routes.IMainStop]
}

type subItinerary struct {
	drone  vehicles.IDrone
	flight routes.ISubRoute
}

type itinerary struct {
	car              vehicles.ICar
	route            routes.IMainRoute
	dronesAndFlights map[DroneNumber]subItinerary
}

func New(car vehicles.ICar) Itinerary {
	return itinerary{
		car:              car,
		dronesAndFlights: generateDronesAndFlights(car.Drones()),
		route:            routes.NewMainRoute(routes.NewMainStop(car.ActualPoint())),
	}
}

func (i itinerary) ActualCarPoint() gps.Point {
	return i.car.ActualPoint()
}

func (i itinerary) CarSupport(nextPoints ...gps.Point) bool {
	return i.car.Support(nextPoints...)
}

func (i itinerary) DroneCanReach(droneNumber DroneNumber, nextPoints ...gps.Point) bool {
	drone := i.dronesAndFlights[droneNumber].drone
	return drone.CanReach(nextPoints...)
}

func (i itinerary) DroneNumbers() []DroneNumber {
	var droneNumbers []DroneNumber
	for drone := range i.dronesAndFlights {
		droneNumbers = append(droneNumbers, drone)
	}
	return droneNumbers
}

func (i itinerary) DroneIsFlying(droneNumber DroneNumber) bool {
	subItn := i.dronesAndFlights[droneNumber]
	return subItn.flight != nil
}

func (i itinerary) DroneSupport(droneNumber DroneNumber, nextPoints ...gps.Point) bool {
	subItn := i.dronesAndFlights[droneNumber]
	return subItn.drone.Support(nextPoints...)
}

func (i itinerary) LandDrone(droneNumber DroneNumber, destination routes.IMainStop) {
	subItn := i.dronesAndFlights[droneNumber]
	if flight := subItn.flight; flight != nil {
		flight.Return(destination)
		subItn.drone.Land(destination.Point())
		subItn.flight = nil
		i.dronesAndFlights[droneNumber] = subItn
	}
}

func (i itinerary) LandAllDrones(landingStop routes.IMainStop) {
	for droneNumber, subItn := range i.dronesAndFlights {
		if flight := subItn.flight; flight != nil {
			flight.Return(landingStop)
			subItn.drone.Land(landingStop.Point())
			subItn.flight = nil
			i.dronesAndFlights[droneNumber] = subItn
		}
	}
}

func (i itinerary) MoveCar(destination gps.Point) {
	i.route.Append(routes.NewMainStop(destination))
	i.car.Move(destination)
}

func (i itinerary) MoveDrone(droneNumber DroneNumber, destination gps.Point) {
	subItn := i.dronesAndFlights[droneNumber]
	if flight := subItn.flight; flight != nil {
		flight.Append(routes.NewSubStop(destination))
		subItn.drone.Move(destination)
		return
	}

	subItn.flight = flightFactory(i.route.Last())
	subItn.flight.Append(routes.NewSubStop(destination))
	subItn.drone.Move(destination)
	i.dronesAndFlights[droneNumber] = subItn
}

func (i itinerary) RemoveMainStopFromRoute(index int) {
	i.route.RemoveMainStop(index)
}

func (i itinerary) RouteIterator() slc.Iterator[routes.IMainStop] {
	return i.route.Iterator()
}

func generateDronesAndFlights(drones []vehicles.IDrone) map[DroneNumber]subItinerary {
	dronesAndFlights := make(map[DroneNumber]subItinerary)
	for i, drone := range drones {
		dronesAndFlights[DroneNumber(i+1)] = subItinerary{
			drone: drone,
		}
	}
	return dronesAndFlights
}
