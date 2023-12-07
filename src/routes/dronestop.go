package routes

import (
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

type IDroneStop interface {
	Point() *gps.Point
	Drone() vehicles.IDrone
	Flight() *flight
}

type droneStop struct {
	point  *gps.Point
	drone  vehicles.IDrone
	flight *flight
}

func NewDroneStop(point *gps.Point, drone vehicles.IDrone, flight *flight) IDroneStop {
	return &droneStop{
		point:  point,
		drone:  drone,
		flight: flight,
	}
}

func (ds *droneStop) Point() *gps.Point {
	return ds.point
}

func (ds *droneStop) Drone() vehicles.IDrone {
	return ds.drone
}

func (ds *droneStop) Flight() *flight {
	return ds.flight
}
