package itinerary

import (
	"github.com/victorguarana/vehicle-routing/src/route"
	"github.com/victorguarana/vehicle-routing/src/vehicle"
)

var flightFactory = route.NewSubRoute

// TODO: Hide its implementation from other packages tests
// Avoid 'var mockedDrone1 = itinerary.DroneNumber(1)'
type DroneNumber int

//go:generate mockgen -source=itinerary.go -destination=mock/itinerarymock.go
type Itinerary interface {
	Info() Info
	Constructor() Constructor
	Finder() Finder
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
	return &info{i}
}

func (i *itinerary) Constructor() Constructor {
	return constructor{info: &info{i}}
}

func (i *itinerary) Finder() Finder {
	return finder{info: &info{i}}
}

func (i *itinerary) Modifier() Modifier {
	return modifier{info: &info{i}}
}

func generateDroneNumbersMap(drones []vehicle.IDrone) map[DroneNumber]vehicle.IDrone {
	activeSubitineraryMap := make(map[DroneNumber]vehicle.IDrone)
	for i, drone := range drones {
		activeSubitineraryMap[DroneNumber(i+1)] = drone
	}
	return activeSubitineraryMap
}
