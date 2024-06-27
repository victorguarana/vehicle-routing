package itinerary

import (
	"log"

	"github.com/victorguarana/vehicle-routing/internal/gps"
	"github.com/victorguarana/vehicle-routing/internal/route"
)

//go:generate mockgen -source=constructor.go -destination=mock/constructormock.go
type Constructor interface {
	Info
	LandAllDrones(landingStop route.IMainStop)
	LandDrone(droneNumber DroneNumber, destination route.IMainStop)
	MoveCar(destination gps.Point)
	MoveDrone(droneNumber DroneNumber, destination gps.Point)
	StartDroneFlight(droneNumber DroneNumber, startingPoint route.IMainStop)
}

type constructor struct {
	*info
}

func (c constructor) LandDrone(droneNumber DroneNumber, destination route.IMainStop) {
	flight := c.activeFlightByNumber(droneNumber)
	if flight == nil {
		log.Panic("Drone is not flying")
	}

	flight.Return(destination)
	drone := c.droneByNumber(droneNumber)
	drone.Land(destination.Point())
	c.archiveFlight(droneNumber, flight)
}

func (c constructor) LandAllDrones(landingStop route.IMainStop) {
	for droneNumber, flight := range c.activeFlights {
		if flight == nil {
			continue
		}

		flight.Return(landingStop)
		drone := c.droneByNumber(droneNumber)
		drone.Land(landingStop.Point())
		c.archiveFlight(droneNumber, flight)
	}
}

func (c constructor) MoveCar(destination gps.Point) {
	c.route.Append(route.NewMainStop(destination))
	c.car.Move(destination)
}

func (c constructor) MoveDrone(droneNumber DroneNumber, destination gps.Point) {
	flight := c.activeFlightByNumber(droneNumber)
	if flight == nil {
		log.Panic("Drone is not flying")
	}

	flight.Append(route.NewSubStop(destination))
	drone := c.droneByNumber(droneNumber)
	drone.Move(destination)
}

func (c constructor) StartDroneFlight(droneNumber DroneNumber, startingPoint route.IMainStop) {
	drone := c.droneByNumber(droneNumber)
	drone.TakeOff()
	flight := flightFactory(startingPoint)
	c.saveActiveFlight(droneNumber, flight)
}

func (c constructor) activeFlightByNumber(droneNumber DroneNumber) route.ISubRoute {
	return c.activeFlights[droneNumber]
}

func (c constructor) archiveFlight(droneNumber DroneNumber, flight route.ISubRoute) {
	c.saveActiveFlight(droneNumber, nil)
	subItn := SubItinerary{
		Drone:  c.droneByNumber(droneNumber),
		Flight: flight,
	}
	c.completedSubItineraryList = append(c.completedSubItineraryList, subItn)
}

func (c constructor) saveActiveFlight(droneNumber DroneNumber, flight route.ISubRoute) {
	c.activeFlights[droneNumber] = flight
}
