package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

type IFlight interface {
	TakeoffPoint() ICarStop
	LandingPoint() ICarStop
	// DroneStops() []*droneStop
	Drone() vehicles.IDrone
	AppendDroneStop(IDroneStop)
	Land(ICarStop)
}

type flight struct {
	takeoffPoint ICarStop
	landingPoint ICarStop

	droneStops []IDroneStop
	drone      vehicles.IDrone
}

func NewFlight(takeoffPoint, landingPoint ICarStop, drone vehicles.IDrone) IFlight {
	return &flight{
		takeoffPoint: takeoffPoint,
		landingPoint: landingPoint,
		drone:        drone,
	}
}

func (f *flight) TakeoffPoint() ICarStop {
	return f.takeoffPoint
}

func (f *flight) LandingPoint() ICarStop {
	return f.landingPoint
}

// func (f *flight) DroneStops() []*droneStop {
// 	return f.droneStops
// }

func (f *flight) Drone() vehicles.IDrone {
	return f.drone
}

func (f *flight) AppendDroneStop(ds IDroneStop) {
	f.droneStops = append(f.droneStops, ds)
}

func (f *flight) Land(cs ICarStop) {
	f.landingPoint = cs
	f.drone.Dock()
}
