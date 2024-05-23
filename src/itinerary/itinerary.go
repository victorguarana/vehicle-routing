package itinerary

import (
	"log"

	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/route"
	"github.com/victorguarana/vehicle-routing/src/slc"
	"github.com/victorguarana/vehicle-routing/src/vehicles"
)

var flightFactory = route.NewSubRoute

// TODO: Hide its implementation from other packages tests
// Avoid 'var mockedDrone1 = itinerary.DroneNumber(1)'
type DroneNumber int

type Itinerary interface {
	ActualCarPoint() gps.Point
	ActualCarStop() route.IMainStop
	CarEfficiency() float64
	CarSpeed() float64
	CarSupport(nextPoints ...gps.Point) bool
	DroneCanReach(droneNumber DroneNumber, nextPoints ...gps.Point) bool
	DroneEfficiency() float64
	DroneIsFlying(droneNumber DroneNumber) bool
	DroneNumbers() []DroneNumber
	DroneSpeed() float64
	DroneSupport(droneNumber DroneNumber, deliveryPoint gps.Point, landingPoint gps.Point) bool
	LandAllDrones(landingStop route.IMainStop)
	LandDrone(droneNumber DroneNumber, destination route.IMainStop)
	MoveCar(destination gps.Point)
	MoveDrone(droneNumber DroneNumber, destination gps.Point)
	RemoveMainStopFromRoute(index int)
	RouteIterator() slc.Iterator[route.IMainStop]
	StartDroneFlight(droneNumber DroneNumber, startingPoint route.IMainStop)
}

type subItinerary struct {
	drone  vehicles.IDrone
	flight route.ISubRoute
}

type itinerary struct {
	activeFlights             map[DroneNumber]route.ISubRoute
	car                       vehicles.ICar
	completedSubItineraryList []subItinerary
	droneNumbersMap           map[DroneNumber]vehicles.IDrone
	route                     route.IMainRoute
}

func New(car vehicles.ICar) Itinerary {
	return &itinerary{
		activeFlights:             map[DroneNumber]route.ISubRoute{},
		car:                       car,
		completedSubItineraryList: []subItinerary{},
		droneNumbersMap:           generateDroneNumbersMap(car.Drones()),
		route:                     route.NewMainRoute(route.NewMainStop(car.ActualPoint())),
	}
}

func (i *itinerary) ActualCarPoint() gps.Point {
	return i.car.ActualPoint()
}

func (i *itinerary) ActualCarStop() route.IMainStop {
	return i.route.Last()
}

func (i *itinerary) CarEfficiency() float64 {
	return vehicles.CarEfficiency
}

func (i *itinerary) CarSpeed() float64 {
	return vehicles.CarSpeed
}

func (i *itinerary) CarSupport(nextPoints ...gps.Point) bool {
	return i.car.Support(nextPoints...)
}

func (i *itinerary) DroneCanReach(droneNumber DroneNumber, nextPoints ...gps.Point) bool {
	drone := i.droneByNumber(droneNumber)
	return drone.CanReach(nextPoints...)
}

func (i *itinerary) DroneEfficiency() float64 {
	return vehicles.DroneEfficiency
}

func (i *itinerary) DroneIsFlying(droneNumber DroneNumber) bool {
	drone := i.droneByNumber(droneNumber)
	return drone.IsFlying()
}

func (i *itinerary) DroneNumbers() []DroneNumber {
	var droneNumbers []DroneNumber
	for droneNumber := range i.droneNumbersMap {
		droneNumbers = append(droneNumbers, droneNumber)
	}
	return droneNumbers
}

func (i *itinerary) DroneSpeed() float64 {
	return vehicles.DroneSpeed
}

func (i *itinerary) DroneSupport(droneNumber DroneNumber, deliveryPoint gps.Point, landingPoint gps.Point) bool {
	drone := i.droneByNumber(droneNumber)
	return drone.Support(deliveryPoint) && drone.CanReach(deliveryPoint, landingPoint)
}

func (i *itinerary) StartDroneFlight(droneNumber DroneNumber, startingPoint route.IMainStop) {
	drone := i.droneByNumber(droneNumber)
	drone.TakeOff()
	flight := flightFactory(startingPoint)
	i.saveActiveFlight(droneNumber, flight)
}

func (i *itinerary) LandDrone(droneNumber DroneNumber, destination route.IMainStop) {
	flight := i.activeFlightByNumber(droneNumber)
	if flight == nil {
		log.Panic("Drone is not flying")
	}

	flight.Return(destination)
	drone := i.droneByNumber(droneNumber)
	drone.Land(destination.Point())
	i.achiveFlight(droneNumber, flight)
}

func (i *itinerary) LandAllDrones(landingStop route.IMainStop) {
	for droneNumber, flight := range i.activeFlights {
		if flight == nil {
			continue
		}

		flight.Return(landingStop)
		drone := i.droneByNumber(droneNumber)
		drone.Land(landingStop.Point())
		i.achiveFlight(droneNumber, flight)
	}
}

func (i *itinerary) MoveCar(destination gps.Point) {
	i.route.Append(route.NewMainStop(destination))
	i.car.Move(destination)
}

func (i *itinerary) MoveDrone(droneNumber DroneNumber, destination gps.Point) {
	flight := i.activeFlightByNumber(droneNumber)
	if flight == nil {
		log.Panic("Drone is not flying")
	}

	flight.Append(route.NewSubStop(destination))
	drone := i.droneByNumber(droneNumber)
	drone.Move(destination)
}

func (i *itinerary) RemoveMainStopFromRoute(index int) {
	i.route.RemoveMainStop(index)
}

func (i *itinerary) RouteIterator() slc.Iterator[route.IMainStop] {
	return i.route.Iterator()
}

func (i *itinerary) achiveFlight(droneNumber DroneNumber, flight route.ISubRoute) {
	i.saveActiveFlight(droneNumber, nil)
	subItn := subItinerary{
		drone:  i.droneByNumber(droneNumber),
		flight: flight,
	}
	i.completedSubItineraryList = append(i.completedSubItineraryList, subItn)
}

func (i *itinerary) droneByNumber(droneNumber DroneNumber) vehicles.IDrone {
	return i.droneNumbersMap[droneNumber]
}

func (i *itinerary) activeFlightByNumber(droneNumber DroneNumber) route.ISubRoute {
	return i.activeFlights[droneNumber]
}

func (i *itinerary) saveActiveFlight(droneNumber DroneNumber, flight route.ISubRoute) {
	i.activeFlights[droneNumber] = flight
}

func generateDroneNumbersMap(drones []vehicles.IDrone) map[DroneNumber]vehicles.IDrone {
	activeSubitineraryMap := make(map[DroneNumber]vehicles.IDrone)
	for i, drone := range drones {
		activeSubitineraryMap[DroneNumber(i+1)] = drone
	}
	return activeSubitineraryMap
}
