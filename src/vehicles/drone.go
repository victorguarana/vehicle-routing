package vehicles

import (
	"log"

	"github.com/victorguarana/vehicle-routing/src/gps"
)

const DroneRange = 150.0
const DroneSpeed = 25.0
const DroneStorage = 10.0
const DroneEfficiency = 30.0

type IDrone interface {
	ActualPoint() gps.Point
	CanReach(...gps.Point) bool
	Efficiency() float64
	IsFlying() bool
	Land(destination gps.Point)
	Move(destination gps.Point)
	Name() string
	Speed() float64
	Support(...gps.Point) bool
	TakeOff()
}

type drone struct {
	actualPoint     gps.Point
	car             *car
	efficiency      float64
	name            string
	speed           float64
	isFlying        bool
	remaningRange   float64
	remaningStorage float64
	totalRange      float64
	totalStorage    float64
}

func newDrone(name string, c *car) *drone {
	return &drone{
		car:             c,
		efficiency:      DroneEfficiency,
		name:            name,
		speed:           DroneSpeed,
		remaningRange:   DroneRange,
		remaningStorage: DroneStorage,
		totalRange:      DroneRange,
		totalStorage:    DroneStorage,
	}
}

func (d *drone) ActualPoint() gps.Point {
	return d.actualPoint
}

func (d *drone) CanReach(route ...gps.Point) bool {
	distance := gps.DistanceBetweenPoints(route...)
	return distance <= d.remaningRange
}

func (d *drone) Efficiency() float64 {
	return d.efficiency
}

func (d *drone) IsFlying() bool {
	return d.isFlying
}

func (d *drone) Land(destination gps.Point) {
	d.isFlying = false
	d.actualPoint = destination
	d.resetAttributes()
}

func (d *drone) Move(destination gps.Point) {
	if !d.isFlying {
		log.Printf("Move: Drone %s moving without take off\n", d.name)
	}
	d.remaningRange -= gps.DistanceBetweenPoints(d.actualPoint, destination)
	d.remaningStorage -= destination.PackageSize
	d.actualPoint = destination
}

func (d *drone) Name() string {
	return d.name
}

func (d *drone) Speed() float64 {
	return d.speed
}

func (d *drone) Support(route ...gps.Point) bool {
	distance := gps.DistanceBetweenPoints(route...)
	packagesSize := 0.0
	for _, destination := range route {
		packagesSize += destination.PackageSize
	}
	if distance > d.remaningRange {
		return false
	}
	if packagesSize > d.remaningStorage {
		return false
	}
	return true
}

func (d *drone) resetAttributes() {
	d.remaningRange = d.totalRange
	d.remaningStorage = d.totalStorage
}

func (d *drone) TakeOff() {
	d.isFlying = true
}
