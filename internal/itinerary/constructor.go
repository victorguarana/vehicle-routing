package itinerary

import (
	"log"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
	"github.com/victorguarana/vehicle-routing/internal/vehicle"
)

//go:generate mockgen -source=constructor.go -destination=mock/constructormock.go
type Constructor interface {
	Info
	LandAllDrones(landingStop route.IMainStop)
	LandDrone(drone vehicle.IDrone, destination route.IMainStop)
	MoveCar(destination gps.Point)
	MoveDrone(drone vehicle.IDrone, destination gps.Point)
	StartDroneFlight(drone vehicle.IDrone, startingPoint route.IMainStop)
}

type constructor struct {
	*info
}

func (c constructor) LandDrone(drone vehicle.IDrone, destination route.IMainStop) {
	flight, ok := c.activeFlights[drone]
	if !ok {
		log.Panic("Drone is not flying")
	}

	flight.Return(destination)
	drone.Land(destination.Point())
	c.archiveFlight(drone, flight)
}

func (c constructor) LandAllDrones(landingStop route.IMainStop) {
	for drone, flight := range c.activeFlights {
		if flight == nil {
			continue
		}

		flight.Return(landingStop)
		drone.Land(landingStop.Point())
		c.archiveFlight(drone, flight)
	}
}

func (c constructor) MoveCar(destination gps.Point) {
	c.route.Append(route.NewMainStop(destination))
	c.car.Move(destination)
}

func (c constructor) MoveDrone(drone vehicle.IDrone, destination gps.Point) {
	flight, ok := c.activeFlights[drone]
	if !ok {
		log.Panic("Drone is not flying")
	}

	flight.Append(route.NewSubStop(destination))
	drone.Move(destination)
}

func (c constructor) StartDroneFlight(drone vehicle.IDrone, startingPoint route.IMainStop) {
	drone.TakeOff()
	flight := flightFactory(startingPoint)
	c.saveActiveFlight(drone, flight)
}

func (c constructor) archiveFlight(drone vehicle.IDrone, flight route.ISubRoute) {
	c.saveActiveFlight(drone, nil)
	subItn := SubItinerary{
		Drone:  drone,
		Flight: flight,
	}
	c.completedSubItineraryList = append(c.completedSubItineraryList, subItn)
}

func (c constructor) saveActiveFlight(drone vehicle.IDrone, flight route.ISubRoute) {
	c.activeFlights[drone] = flight
}
