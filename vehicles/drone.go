package vehicles

import "github.com/victorguarana/go-vehicle-route/gps"

type IDrone interface {
	ivehicle
}

type drone struct {
	vehicle
	isFlying        bool
	totalStorage    float64
	remaningStorage float64
	totalRange      float64
	remaningRange   float64
}

func NewDrone(name string, startingPoint *gps.Point) IDrone {
	d := drone{
		totalStorage:    defaultStorage,
		remaningStorage: defaultStorage,
		totalRange:      defaultRange,
		remaningRange:   defaultRange,
		vehicle: vehicle{
			speed:          defaultSpeed,
			name:           name,
			actualPosition: startingPoint,
		},
	}
	return &d
}

func (d *drone) Move(destination *gps.Point) error {
	if d.actualPosition == nil || destination == nil {
		return ErrInvalidParams
	}

	distance := gps.DistanceBetweenPoints(*d.actualPosition, *destination)
	if distance >= d.remaningRange {
		return ErrSoFar
	}

	d.remaningRange -= distance
	d.actualPosition = destination

	return nil
}

func (d *drone) Reachable(destination gps.Point) bool {
	distance := gps.DistanceBetweenPoints(*d.actualPosition, destination)
	return distance >= d.remaningRange
}
