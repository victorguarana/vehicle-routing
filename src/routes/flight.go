package routes

import (
	"errors"

	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

var (
	ErrInvalidTakeoffPoint = errors.New("invalid takeoff point")
	ErrNilDrone            = errors.New("drone can not be nil")
	ErrInvalidDroneStop    = errors.New("invalid drone stop")
)

type IFlight interface {
	TakeoffPoint() ICarStop
	LandingPoint() ICarStop
	Drone() vehicles.IDrone

	Land(ICarStop) error
	Append(*gps.Point) error
}

type flight struct {
	takeoffPoint *carStop
	landingPoint *carStop

	stops []*droneStop
	drone vehicles.IDrone
}

func NewFlight(drone vehicles.IDrone, takeoffPoint, landingPoint ICarStop) (IFlight, error) {
	if drone == nil {
		return nil, ErrNilDrone
	}

	takeoffPointStruct, ok := takeoffPoint.(*carStop)
	if !ok {
		return nil, ErrInvalidTakeoffPoint
	}

	landingPointStruct, _ := landingPoint.(*carStop)

	return &flight{
		takeoffPoint: takeoffPointStruct,
		landingPoint: landingPointStruct,
		drone:        drone,
		stops:        []*droneStop{},
	}, nil
}

func (f *flight) TakeoffPoint() ICarStop {
	return f.takeoffPoint
}

func (f *flight) LandingPoint() ICarStop {
	return f.landingPoint
}

func (f *flight) Drone() vehicles.IDrone {
	return f.drone
}

func (f *flight) Append(point *gps.Point) error {
	ds := newDroneStop(f.drone, point, f)

	f.stops = append(f.stops, ds)
	return nil
}

func (f *flight) Land(cs ICarStop) error {
	var ok bool
	f.landingPoint, ok = cs.(*carStop)
	if !ok {
		return ErrInvalidCarStop
	}

	return nil
}
