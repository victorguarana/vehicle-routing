package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

type IDroneStop interface {
	Drone() vehicles.IDrone
	Flight() IFlight
	Point() gps.Point
}

type droneStop struct {
	point  gps.Point
	drone  vehicles.IDrone
	flight *flight
}

func newDroneStop(d vehicles.IDrone, p gps.Point, f *flight) *droneStop {
	return &droneStop{
		point:  p,
		drone:  d,
		flight: f,
	}
}

func (ds *droneStop) Drone() vehicles.IDrone {
	return ds.drone
}

func (ds *droneStop) Flight() IFlight {
	return ds.flight
}

func (ds *droneStop) Point() gps.Point {
	return ds.point
}
