package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

type Itinerary struct {
	Car   vehicles.ICar
	Route IMainRoute
}

func NewItinerary(car vehicles.ICar) Itinerary {
	initialPosition := NewMainStop(car.ActualPosition()).(*mainStop)
	return Itinerary{
		Car: car,
		Route: &mainRoute{
			mainStops: []*mainStop{initialPosition},
		},
	}
}
