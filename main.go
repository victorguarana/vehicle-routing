package main

import (
	"github.com/victorguarana/go-vehicle-route/src/csp"
	"github.com/victorguarana/go-vehicle-route/src/gps"
	"github.com/victorguarana/go-vehicle-route/src/greedy"
	"github.com/victorguarana/go-vehicle-route/src/itinerary"
	"github.com/victorguarana/go-vehicle-route/src/output"
	"github.com/victorguarana/go-vehicle-route/src/vehicles"
)

func main() {
	ClosestNeighbor()
	ClosestNeighborWithDrones()
	Covering()
}

func ClosestNeighbor() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicles.NewCar("car1", initialPoint)
	itn := itinerary.New(car)
	greedy.ClosestNeighbor([]itinerary.Itinerary{itn}, gps.GetMap())
	filename := "closest-neighbor.png"
	output.ToImage(filename, itn)
}

func ClosestNeighborWithDrones() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicles.NewCar("car1", initialPoint)
	car.NewDrone("drone1")

	itn := itinerary.New(car)
	greedy.ClosestNeighbor([]itinerary.Itinerary{itn}, gps.GetMap())
	greedy.DroneStrikesInsertion(itn)
	filename := "closest-neighbor-with-drones.png"
	output.ToImage(filename, itn)
}

func Covering() {
	initialPoint := gps.Point{Name: "initialPoint"}
	car := vehicles.NewCar("car1", initialPoint)
	car.NewDrone("drone1")
	itn := itinerary.New(car)
	neighorhoodDistance := vehicles.DroneRange / 4
	csp.CoveringWithDrones([]itinerary.Itinerary{itn}, gps.GetMap(), neighorhoodDistance)
	filename := "covering.png"
	output.ToImage(filename, itn)
}
