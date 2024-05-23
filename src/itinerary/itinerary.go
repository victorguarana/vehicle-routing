package itinerary

import (
	"github.com/victorguarana/vehicle-routing/src/gps"
	"github.com/victorguarana/vehicle-routing/src/route"
	"github.com/victorguarana/vehicle-routing/src/slc"
	"github.com/victorguarana/vehicle-routing/src/vehicle"
)

var flightFactory = route.NewSubRoute

// TODO: Hide its implementation from other packages tests
// Avoid 'var mockedDrone1 = itinerary.DroneNumber(1)'
type DroneNumber int

type Itinerary interface {
	Info() Info
	Constructor() Constructor
	Modifier() Modifier
}

type SubItinerary struct {
	Drone  vehicle.IDrone
	Flight route.ISubRoute
}

type itinerary struct {
	activeFlights             map[DroneNumber]route.ISubRoute
	car                       vehicle.ICar
	completedSubItineraryList []SubItinerary
	droneNumbersMap           map[DroneNumber]vehicle.IDrone
	route                     route.IMainRoute
}

func New(car vehicle.ICar) Itinerary {
	return &itinerary{
		activeFlights:             map[DroneNumber]route.ISubRoute{},
		car:                       car,
		completedSubItineraryList: []SubItinerary{},
		droneNumbersMap:           generateDroneNumbersMap(car.Drones()),
		route:                     route.NewMainRoute(route.NewMainStop(car.ActualPoint())),
	}
}

func (i *itinerary) Info() Info {
	return info{i}
}

func (i *itinerary) ActualCarPoint() gps.Point {
	return i.car.ActualPoint()
}

func (i *itinerary) ActualCarStop() route.IMainStop {
	return i.route.Last()
}

func (i *itinerary) CarEfficiency() float64 {
	return vehicle.CarEfficiency
}

func (i *itinerary) CarSpeed() float64 {
	return vehicle.CarSpeed
}

func (i *itinerary) CarSupport(nextPoints ...gps.Point) bool {
	return i.car.Support(nextPoints...)
}

func (i *itinerary) Constructor() Constructor {
	return constructor{info: &info{i}}
}

func (i *itinerary) DroneCanReach(droneNumber DroneNumber, nextPoints ...gps.Point) bool {
	drone := i.droneByNumber(droneNumber)
	return drone.CanReach(nextPoints...)
}

func (i *itinerary) DroneEfficiency() float64 {
	return vehicle.DroneEfficiency
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
	return vehicle.DroneSpeed
}

func (i *itinerary) DroneSupport(droneNumber DroneNumber, deliveryPoint gps.Point, landingPoint gps.Point) bool {
	drone := i.droneByNumber(droneNumber)
	return drone.Support(deliveryPoint) && drone.CanReach(deliveryPoint, landingPoint)
}

func (i *itinerary) Modifier() Modifier {
	return modifier{info: &info{i}}
}

func (i *itinerary) RouteIterator() slc.Iterator[route.IMainStop] {
	return i.route.Iterator()
}

func (i *itinerary) SubItineraryList() []SubItinerary {
	return i.completedSubItineraryList
}

func (i *itinerary) droneByNumber(droneNumber DroneNumber) vehicle.IDrone {
	return i.droneNumbersMap[droneNumber]
}

func generateDroneNumbersMap(drones []vehicle.IDrone) map[DroneNumber]vehicle.IDrone {
	activeSubitineraryMap := make(map[DroneNumber]vehicle.IDrone)
	for i, drone := range drones {
		activeSubitineraryMap[DroneNumber(i+1)] = drone
	}
	return activeSubitineraryMap
}
