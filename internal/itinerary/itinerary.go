package itinerary

import (
	"github.com/victorguarana/vehicle-routing/internal/route"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

type ItineraryList []Itinerary

var flightFactory = route.NewSubRoute

//go:generate mockgen -source=itinerary.go -destination=mock/itinerarymock.go
type Itinerary interface {
	Info() Info
	Constructor() Constructor
	Finder() Finder
	Modifier() Modifier
	Validator() Validator
}

type SubItinerary struct {
	Drone  vehicle.IDrone
	Flight route.ISubRoute
}

type itinerary struct {
	activeFlights             map[vehicle.IDrone]route.ISubRoute
	car                       vehicle.ICar
	completedSubItineraryList []SubItinerary
	route                     route.IMainRoute
}

func New(car vehicle.ICar) Itinerary {
	return &itinerary{
		activeFlights:             map[vehicle.IDrone]route.ISubRoute{},
		car:                       car,
		completedSubItineraryList: []SubItinerary{},
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

func (i *itinerary) Validator() Validator {
	return &validator{info: &info{i}}
}
