package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

type IFlight interface {
	TakeoffPoint() *carStop
	LandingPoint() *carStop
	DroneStops() []*droneStop
	Drone() vehicles.IDrone
	AppendDroneStop(*droneStop)
}

type flight struct {
	takeoffPoint *carStop
	landingPoint *carStop

	droneStops []*droneStop
	drone      vehicles.IDrone
}

func NewFlight(takeoffPoint, landingPoint *carStop, drone vehicles.IDrone) IFlight {
	return &flight{
		takeoffPoint: takeoffPoint,
		landingPoint: landingPoint,
		drone:        drone,
	}
}

func (f *flight) TakeoffPoint() *carStop {
	return f.takeoffPoint
}

func (f *flight) LandingPoint() *carStop {
	return f.landingPoint
}

func (f *flight) DroneStops() []*droneStop {
	return f.droneStops
}

func (f *flight) Drone() vehicles.IDrone {
	return f.drone
}

func (f *flight) AppendDroneStop(ds *droneStop) {
	f.droneStops = append(f.droneStops, ds)
}

func (f *flight) Land(cs *carStop) {
	f.landingPoint = cs
}
